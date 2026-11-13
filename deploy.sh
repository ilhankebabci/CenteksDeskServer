#!/bin/bash
set -e

echo "=========================================================="
echo " 🚀 CENTEKS Destek Yönetim Sunucusu Güncelleniyor..."
echo "=========================================================="

REPO_DIR="/opt/CenteksDeskServer"
DATA_DIR="/opt/rustdesk-api-pro"

# 1. Projeyi çek veya güncelle
if [ ! -d "$REPO_DIR" ]; then
    echo ">>> Kod deposu klonlanıyor..."
    sudo git clone https://github.com/ilhankebabci/CenteksDeskServer.git "$REPO_DIR"
else
    echo ">>> Güncel kodlar çekiliyor..."
    cd "$REPO_DIR"
    sudo git fetch origin master
    sudo git reset --hard origin/master
fi

cd "$REPO_DIR"

# 1.5 İstemci (.exe) dosyasını downloads dizinine hazırla
sudo mkdir -p /var/www/centeks-panel/downloads
if [ ! -f /var/www/centeks-panel/downloads/CenteksDestek-x86_64.exe ]; then
    echo ">>> Güncel Windows istemcisi (.exe) indiriliyor..."
    sudo curl -L https://github.com/ilhankebabci/rustdesk/releases/download/v1.5.3/CenteksDestek-x86_64.exe -o /var/www/centeks-panel/downloads/CenteksDestek-x86_64.exe || true
fi

# 2. Docker imajını derle
echo ">>> Docker imajı derleniyor (centeks-server:latest)..."
sudo docker build -t centeks-server:latest .

# 3. Eski konteyneri durdur
if [ "$(sudo docker ps -aq -f name=rustdesk-api-server-pro)" ]; then
    echo ">>> Eski konteyner durduruluyor..."
    sudo docker stop rustdesk-api-server-pro > /dev/null 2>&1 || true
    sudo docker rm rustdesk-api-server-pro > /dev/null 2>&1 || true
fi

# 4. Yeni CENTEKS imajını başlat
echo ">>> CENTEKS Destek Paneli başlatılıyor..."
sudo docker run -d \
  --name rustdesk-api-server-pro \
  --restart unless-stopped \
  -e ADMIN_USER=admin \
  -e ADMIN_PASS=Centeks2027! \
  -e TZ=Europe/Istanbul \
  -p 8080:8080 \
  -v "$DATA_DIR/server.yaml":/app/server.yaml \
  -v "$DATA_DIR/server.yaml":/app/data/server.yaml \
  -v "$DATA_DIR":/app/data \
  centeks-server:latest

sleep 5

# 5. Yönetici hesabını doğrula
echo ">>> Yönetici hesabı doğrulanıyor..."
sudo docker exec rustdesk-api-server-pro rustdesk-api-server-pro user add admin Centeks2027! --admin > /dev/null 2>&1 || true

echo "=========================================================="
echo " ✅ CENTEKS Destek Paneli Başarıyla Güncellendi!"
echo " 🌐 Web Adresi: https://panel.centeks.com"
echo " 👤 Kullanıcı Adı: admin"
echo " 🔑 Parola: Centeks2027!"
echo "=========================================================="
