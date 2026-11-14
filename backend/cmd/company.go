package cmd

import (
	"fmt"
	"strings"

	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/db"

	"github.com/spf13/cobra"
	"xorm.io/xorm"
)

var companyCmd = &cobra.Command{
	Use:   "company",
	Short: "Company & Department management",
}

var deptFlag string

var companyAddCmd = &cobra.Command{
	Use:   "add [company_name]",
	Short: "Add a company with optional departments, e.g.: company add 'Akeray' -d 'Muhasebe,Üretim,Sevkiyat'",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		companyName := strings.TrimSpace(args[0])
		if companyName == "" {
			fmt.Println("Firma adı boş olamaz")
			return
		}

		cfg := config.GetServerConfig()
		engine, err := db.NewEngine(cfg.Db)
		if err != nil {
			fmt.Println("Db Engine create error:", err)
			return
		}

		var existing model.Company
		has, err := engine.Where("name = ?", companyName).Get(&existing)
		if err != nil {
			fmt.Println("DB error:", err)
			return
		}
		if !has {
			_, err = engine.Insert(&model.Company{Name: companyName})
			if err != nil {
				fmt.Println("Firma eklenemedi:", err)
				return
			}
			fmt.Printf("🏢 Firma başarıyla eklendi: %s\n", companyName)
		} else {
			fmt.Printf("🏢 Firma zaten mevcut: %s\n", companyName)
		}

		if deptFlag != "" {
			depts := strings.Split(deptFlag, ",")
			for _, d := range depts {
				d = strings.TrimSpace(d)
				if d == "" {
					continue
				}
				var exDept model.Department
				hasD, _ := engine.Where("company_name = ? AND name = ?", companyName, d).Get(&exDept)
				if !hasD {
					engine.Insert(&model.Department{
						CompanyName: companyName,
						Name:        d,
					})
					fmt.Printf("  📁 Departman eklendi: %s -> %s\n", companyName, d)
				}
			}
		}
	},
}

var deptAddCmd = &cobra.Command{
	Use:   "add-dept [company_name] [department_name]",
	Short: "Add a department to a company, e.g.: company add-dept 'Centeks' 'Bilgi İşlem'",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		companyName := strings.TrimSpace(args[0])
		deptName := strings.TrimSpace(args[1])
		if companyName == "" || deptName == "" {
			fmt.Println("Firma adı ve departman adı boş olamaz")
			return
		}

		cfg := config.GetServerConfig()
		engine, err := db.NewEngine(cfg.Db)
		if err != nil {
			fmt.Println("Db Engine create error:", err)
			return
		}

		var comp model.Company
		hasC, _ := engine.Where("name = ?", companyName).Get(&comp)
		if !hasC {
			engine.Insert(&model.Company{Name: companyName})
			fmt.Printf("🏢 Firma oluşturuldu: %s\n", companyName)
		}

		var exDept model.Department
		hasD, _ := engine.Where("company_name = ? AND name = ?", companyName, deptName).Get(&exDept)
		if !hasD {
			engine.Insert(&model.Department{
				CompanyName: companyName,
				Name:        deptName,
			})
			fmt.Printf("📁 Departman eklendi: %s -> %s\n", companyName, deptName)
		} else {
			fmt.Printf("📁 Departman zaten mevcut: %s -> %s\n", companyName, deptName)
		}
	},
}

var companyListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all companies and their departments",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.GetServerConfig()
		engine, err := db.NewEngine(cfg.Db)
		if err != nil {
			fmt.Println("Db Engine create error:", err)
			return
		}

		var companies []model.Company
		engine.Find(&companies)

		var departments []model.Department
		engine.Find(&departments)

		deptMap := map[string][]string{}
		for _, d := range departments {
			deptMap[d.CompanyName] = append(deptMap[d.CompanyName], d.Name)
		}

		fmt.Println("=== CENTEKS KURUMSAL ŞİRKET & DEPARTMAN LİSTESİ ===")
		for _, c := range companies {
			count, _ := engine.Where("company = ?", c.Name).Count(&model.Device{})
			fmt.Printf("🏢 %s (Cihaz Sayısı: %d)\n", c.Name, count)
			for _, d := range deptMap[c.Name] {
				fmt.Printf("   └── 📁 %s\n", d)
			}
		}
	},
}

// SeedDefaultCompanies varsayılan Centeks ve Akeray şirketlerini ve departmanlarını otomatik ekler
func SeedDefaultCompanies(engine *xorm.Engine) {
	defaults := map[string][]string{
		"Centeks Group": {"Bilgi İşlem", "Muhasebe", "Kalite", "Fabrika", "Sunucular"},
		"Akeray":        {"Muhasebe", "Üretim", "Sevkiyat", "Bilgi İşlem"},
	}

	for compName, depts := range defaults {
		var comp model.Company
		has, _ := engine.Where("name = ?", compName).Get(&comp)
		if !has {
			engine.Insert(&model.Company{Name: compName})
		}
		for _, d := range depts {
			var dept model.Department
			hasD, _ := engine.Where("company_name = ? AND name = ?", compName, d).Get(&dept)
			if !hasD {
				engine.Insert(&model.Department{
					CompanyName: compName,
					Name:        d,
				})
			}
		}
	}
}

func init() {
	companyAddCmd.Flags().StringVarP(&deptFlag, "departments", "d", "", "Virgülle ayrılmış departmanlar (ör: 'Muhasebe,Kalite,Fabrika')")
	companyCmd.AddCommand(companyAddCmd)
	companyCmd.AddCommand(deptAddCmd)
	companyCmd.AddCommand(companyListCmd)
	RootCmd.AddCommand(companyCmd)
}
