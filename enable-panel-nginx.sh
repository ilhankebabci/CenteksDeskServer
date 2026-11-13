#!/bin/bash
set -e

echo "=========================================================="
echo " 🚀 CENTEKS Pro Paneli Varsayılan Yapılıyor (Port 8080 -> 443)..."
echo "=========================================================="

# 1. Eski köprü servisini durdur
echo ">>> Eski köprü servisi (centeks-bridge) durduruluyor..."
sudo systemctl stop centeks-bridge > /dev/null 2>&1 || true
sudo systemctl disable centeks-bridge > /dev/null 2>&1 || true

# 2. Nginx ayarını güncelle
NGINX_CONF="/etc/nginx/sites-available/panel.centeks.com"

if [ -f "$NGINX_CONF" ]; then
    echo ">>> Nginx konfigürasyonu güncelleniyor ($NGINX_CONF)..."
    sudo cp "$NGINX_CONF" "${NGINX_CONF}.bak"
    cat << 'EOF' | sudo tee "$NGINX_CONF" > /dev/null
# HTTP -> HTTPS Yönlendirmesi
server {
    listen 80;
    listen [::]:80;
    server_name panel.centeks.com 192.168.16.19;
    return 301 https://$host$request_uri;
}

# HTTPS - CENTEKS Destek Pro Yönetim Paneli & API
server {
    listen 443 ssl;
    listen [::]:443 ssl;
    server_name panel.centeks.com 192.168.16.19;

    # SSL Sertifikaları
    ssl_certificate /etc/letsencrypt/live/panel.centeks.com/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/panel.centeks.com/privkey.pem;
    include /etc/letsencrypt/options-ssl-nginx.conf;
    ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem;

    # Güvenlik Başlıkları
    add_header X-Frame-Options "SAMEORIGIN" always;
    add_header X-XSS-Protection "1; mode=block" always;
    add_header X-Content-Type-Options "nosniff" always;
    add_header Referrer-Policy "strict-origin-when-cross-origin" always;

    # CENTEKS Destek Pro Yönetim Paneli & API (Port 8080)
    location / {
        proxy_pass http://127.0.0.1:8080;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_read_timeout 86400s;
        proxy_send_timeout 86400s;
    }

    # İstemci (.exe) İndirme Dizini
    location /downloads/ {
        alias /var/www/centeks-panel/downloads/;
        autoindex off;
    }
}
EOF

    sudo nginx -t
    sudo systemctl reload nginx
    echo "✅ Nginx başarıyla güncellendi ve yenilendi!"
else
    echo "⚠️ $NGINX_CONF bulunamadı!"
fi

echo "=========================================================="
echo " ✅ Tamamlandı! Artık doğrudan şu adresten erişebilirsin:"
echo " 🌐 https://panel.centeks.com (Port yazmaya gerek kalmadı!)"
echo "=========================================================="

