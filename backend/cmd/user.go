package cmd

import (
	"fmt"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/db"
	"rustdesk-api-server-pro/util"

	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "User management",
}

var isAdmin bool

var userAddCmd = &cobra.Command{
	Use:   "add",
	Short: "add user [ add username password][--admin]",
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		password, _ := util.Password(args[1])

		cfg := config.GetServerConfig()
		engine, err := db.NewEngine(cfg.Db)
		if err != nil {
			fmt.Println("Db Engine create error:", err)
			return
		}

		// Kullanıcı zaten var mı kontrol et
		var existing model.User
		has, err := engine.Where("username = ?", username).Asc("id").Get(&existing)
		if err == nil && has {
			// Kullanıcı zaten var, parolasını ve yetkisini güncelle
			existing.Password = password
			existing.IsAdmin = isAdmin
			existing.Status = 1
			engine.Where("id = ?", existing.Id).Cols("password", "is_admin", "status").Update(&existing)

			// Varsa aynı isimli mükerrer fazlalık kayıtları otomatik temizle
			engine.Where("username = ? AND id != ?", username, existing.Id).Delete(&model.User{})

			fmt.Println("User already exists, updated successfully and cleaned duplicates")
			return
		}

		user := &model.User{
			Username:        username,
			Password:        password,
			Name:            username,
			LicensedDevices: 0,
			LoginVerify:     model.LOGIN_ACCESS_TOKEN,
			IsAdmin:         isAdmin,
			Status:          1,
		}

		_, err = engine.Insert(user)
		if err != nil {
			fmt.Println("Add error:", err)
			return
		}
		fmt.Println("Add success")
	},
}

func init() {
	userAddCmd.Flags().BoolVarP(&isAdmin, "admin", "a", false, "Set user admin")
	userCmd.AddCommand(userAddCmd)
	RootCmd.AddCommand(userCmd)
}
