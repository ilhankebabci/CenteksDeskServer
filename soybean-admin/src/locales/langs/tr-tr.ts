const local: App.I18n.Schema = {
  system: {
    title: 'CENTEKS Destek Paneli',
    updateTitle: 'Sistem Güncelleme Bildirimi',
    updateContent: 'Yeni bir sistem sürümü algılandı. Sayfayı hemen yenilemek istiyor musunuz?',
    updateConfirm: 'Hemen Yenile',
    updateCancel: 'Daha Sonra'
  },
  common: {
    action: 'İşlem',
    add: 'Ekle',
    addSuccess: 'Başarıyla Eklendi',
    backToHome: 'Ana Sayfaya Dön',
    batchDelete: 'Toplu Sil',
    cancel: 'İptal',
    close: 'Kapat',
    check: 'Kontrol Et',
    expandColumn: 'Sütunu Genişlet',
    columnSetting: 'Sütun Ayarları',
    config: 'Yapılandırma',
    confirm: 'Onayla',
    delete: 'Sil',
    deleteSuccess: 'Başarıyla Silindi',
    confirmDelete: 'Silmek istediğinizden emin misiniz?',
    edit: 'Düzenle',
    look: 'Görüntüle',
    warning: 'Uyarı',
    error: 'Hata',
    index: 'Sıra',
    keywordSearch: 'Lütfen arama terimi girin',
    logout: 'Çıkış Yap',
    logoutConfirm: 'Çıkış yapmak istediğinizden emin misiniz?',
    lookForward: 'Çok Yakında',
    modify: 'Değiştir',
    modifySuccess: 'Başarıyla Güncellendi',
    noData: 'Veri Bulunamadı',
    operate: 'İşlem',
    pleaseCheckValue: 'Lütfen girilen değerin geçerli olduğunu kontrol edin',
    refresh: 'Yenile',
    reset: 'Sıfırla',
    search: 'Ara',
    switch: 'Değiştir',
    tip: 'İpucu',
    trigger: 'Tetikle',
    update: 'Güncelle',
    updateSuccess: 'Başarıyla Güncellendi',
    userCenter: 'Kullanıcı Profili',
    yesOrNo: {
      yes: 'Evet',
      no: 'Hayır'
    }
  },
  request: {
    logout: 'İstek başarısız olduktan sonra kullanıcı oturumunu kapat',
    logoutMsg: 'Oturum süresi doldu, lütfen tekrar giriş yapın',
    logoutWithModal: 'İstek hatasında açılır pencere ile çıkış yap',
    logoutWithModalMsg: 'Oturum süresi doldu, lütfen tekrar giriş yapın',
    refreshToken: 'Oturum anahtarı yenileniyor',
    tokenExpired: 'Oturum anahtarının süresi doldu'
  },
  theme: {
    themeSchema: {
      title: 'Tema Görünümü',
      light: 'Açık',
      dark: 'Koyu',
      auto: 'Sistemle Eşle'
    },
    grayscale: 'Gri Tonlama',
    colourWeakness: 'Renk Körlüğü Modu',
    layoutMode: {
      title: 'Düzen Modu',
      vertical: 'Dikey Menü Modu',
      horizontal: 'Yatay Menü Modu',
      'vertical-mix': 'Dikey Karma Menü',
      'horizontal-mix': 'Yatay Karma Menü',
      reverseHorizontalMix: 'Birinci ve alt seviye menülerin konumunu ters çevir'
    },
    recommendColor: 'Önerilen Renk Algoritmasını Uygula',
    recommendColorDesc: 'Önerilen renk algoritması referansı',
    themeColor: {
      title: 'Tema Rengi',
      primary: 'Birincil',
      info: 'Bilgi',
      success: 'Başarılı',
      warning: 'Uyarı',
      error: 'Hata',
      followPrimary: 'Birincili Takip Et'
    },
    scrollMode: {
      title: 'Kaydırma Modu',
      wrapper: 'Kapsayıcı',
      content: 'İçerik'
    },
    page: {
      animate: 'Sayfa Animasyonu',
      mode: {
        title: 'Animasyon Türü',
        fade: 'Solma (Fade)',
        'fade-slide': 'Kayma (Slide)',
        'fade-bottom': 'Aşağıdan Solma',
        'fade-scale': 'Ölçeklenerek Solma',
        'zoom-fade': 'Yakınlaşma',
        'zoom-out': 'Uzaklaşma',
        none: 'Yok'
      }
    },
    fixedHeaderAndTab: 'Sabit Başlık ve Sekmeler',
    header: {
      height: 'Başlık Yüksekliği',
      breadcrumb: {
        visible: 'Gezinme Yolu Görünür',
        showIcon: 'Gezinme İkonu Görünür'
      }
    },
    tab: {
      visible: 'Sekmeler Görünür',
      cache: 'Sekme Önbelleği',
      height: 'Sekme Yüksekliği',
      mode: {
        title: 'Sekme Stili',
        chrome: 'Chrome Stili',
        button: 'Buton Stili'
      }
    },
    sider: {
      inverted: 'Koyu Kenar Çubuğu',
      width: 'Kenar Çubuğu Genişliği',
      collapsedWidth: 'Daraltılmış Genişlik',
      mixWidth: 'Karma Genişlik',
      mixCollapsedWidth: 'Karma Daraltılmış Genişlik',
      mixChildMenuWidth: 'Alt Menü Genişliği'
    },
    footer: {
      visible: 'Alt Bilgi Görünür',
      fixed: 'Sabit Alt Bilgi',
      height: 'Alt Bilgi Yüksekliği',
      right: 'Sağ Alt Bilgi'
    },
    watermark: {
      visible: 'Filigran Görünür',
      text: 'Filigran Metni'
    },
    themeDrawerTitle: 'Görünüm Ayarları',
    pageFunTitle: 'Sayfa İşlevleri',
    configOperation: {
      copyConfig: 'Yapılandırmayı Kopyala',
      copySuccessMsg: 'Yapılandırma kopyalandı',
      resetConfig: 'Yapılandırmayı Sıfırla',
      resetSuccessMsg: 'Yapılandırma sıfırlandı'
    }
  },
  route: {
    login: 'Giriş Yap',
    403: 'Yetkisiz Erişim',
    404: 'Sayfa Bulunamadı',
    500: 'Sunucu Hatası',
    'iframe-page': 'Çerçeve Sayfa',
    home: 'Kontrol Paneli',
    audit: 'Oturum & Güvenlik',
    user: 'Kullanıcı Yönetimi',
    user_list: 'Kullanıcı Listesi',
    user_sessions: 'Aktif Oturumlar',
    system: 'Sistem Yönetimi',
    system_mail_template: 'E-posta Şablonları',
    system_mail_logs: 'E-posta Günlükleri',
    system_mail: 'E-posta Yönetimi',
    audit_baselogs: 'Bağlantı Kayıtları',
    audit_filetransferlogs: 'Dosya Aktarım Günlükleri',
    devices: 'Cihaz Yönetimi'
  },
  page: {
    login: {
      common: {
        loginOrRegister: 'Giriş Yap',
        userNamePlaceholder: 'Lütfen kullanıcı adınızı girin',
        phonePlaceholder: 'Lütfen telefon numaranızı girin',
        codePlaceholder: 'Lütfen doğrulama kodunu girin',
        passwordPlaceholder: 'Lütfen parolanızı girin',
        confirmPasswordPlaceholder: 'Lütfen parolanızı tekrar girin',
        codeLogin: 'Kod ile Giriş',
        confirm: 'Giriş Yap',
        back: 'Geri',
        validateSuccess: 'Doğrulama Başarılı',
        loginSuccess: 'Giriş Başarılı',
        welcomeBack: 'Hoş geldiniz, {userName}!'
      },
      pwdLogin: {
        title: 'CENTEKS Giriş Portalı',
        rememberMe: 'Beni Hatırla'
      }
    },
    home: {
      greeting: 'İyi çalışmalar {userName}, CENTEKS operasyon paneli hazır!',
      friendlySponsorship: 'Bilgi Merkezi',
      cupOfCoffee: 'CENTEKS Destek',
      thankYou: 'Aktif Durum',
      userCount: 'Kullanıcı Sayısı',
      deviceCount: 'Toplam Cihaz',
      onlineCount: 'Çevrimiçi Cihaz',
      visitsCount: 'Bağlantı Sayısı',
      operatingSystem: 'İşletim Sistemleri',
      oneWeek: 'Haftalık Aktivite',
      changeLogs: 'Sistem Günlükleri'
    },
    user: {
      list: {
        addUser: 'Kullanıcı Ekle',
        editUser: 'Kullanıcıyı Düzenle',
        inputUsername: 'Kullanıcı Adı Girin',
        inputPassword: 'Parola Girin',
        inputNickname: 'Ad Soyad / Unvan',
        emailFormatError: 'E-posta biçimi geçersiz',
        selectUserStatus: 'Kullanıcı durumunu seçin',
        searchPlaceholder: 'Kullanıcı Adı / Ad / E-posta',
        tfa_secret_bind: '2FA Cihazı Eşleştir',
        require2FASecret: '2FA Anahtarı Boş Olamaz',
        require2FACode: '2FA Kodu Boş Olamaz'
      },
      sessions: {
        kill: 'Oturumu Kapat',
        confirmKill: 'Bu oturumu sonlandırmak istediğinizden emin misiniz?'
      },
      audit: {
        logsSearchPlaceholder: 'Kullanıcı / İşlem / RustDesk ID / IP'
      },
      devices: {
        logsSearchPlaceholder: 'Bilgisayar Adı / Kullanıcı / RustDesk ID'
      }
    },
    system: {
      mailTemplate: {
        addMailTemplate: 'Şablon Ekle',
        editMailTemplate: 'Şablonu Düzenle',
        inputName: 'Şablon Adı',
        inputSubject: 'Konu',
        inputContents: 'İçerik',
        selectType: 'Lütfen tür seçin'
      },
      mailLog: {
        info: 'Detay'
      }
    }
  },
  dropdown: {
    closeCurrent: 'Mevcut Sekmeyi Kapat',
    closeOther: 'Diğerlerini Kapat',
    closeLeft: 'Soldakileri Kapat',
    closeRight: 'Sağdakileri Kapat',
    closeAll: 'Tümünü Kapat'
  },
  icon: {
    themeConfig: 'Görünüm Ayarları',
    themeSchema: 'Tema Değiştir',
    lang: 'Dili Değiştir',
    fullscreen: 'Tam Ekran',
    fullscreenExit: 'Tam Ekrandan Çık',
    reload: 'Sayfayı Yenile',
    collapse: 'Menüyü Daralt',
    expand: 'Menüyü Genişlet',
    pin: 'Sabitle',
    unpin: 'Sabitlemeyi Kaldır'
  },
  datatable: {
    itemCount: 'Toplam {total} kayıt'
  },
  dataMap: {
    user: {
      username: 'Kullanıcı Adı',
      password: 'Parola',
      name: 'Ad Soyad',
      email: 'E-posta',
      licensed_devices: 'Lisanslı Cihazlar',
      login_verify: 'Giriş Doğrulama',
      status: 'Durum',
      is_admin: 'Yönetici mi?',
      tfa_secret: '2FA Gizli Anahtarı',
      tfa_code: '2FA Kodu',
      created_at: 'Oluşturulma Tarihi',
      statusLabel: {
        disabled: 'Devre Dışı',
        unverified: 'Doğrulanmamış',
        normal: 'Aktif'
      },
      loginVerifyLabel: {
        none: 'Doğrulama Yok',
        emailCheck: 'E-posta Doğrulaması',
        tfaCheck: '2FA İki Adımlı Doğrulama'
      }
    },
    session: {
      expired: 'Bitiş Tarihi',
      created_at: 'Oturum Başlangıcı'
    },
    device: {
      username: 'Kullanıcı Adı',
      hostname: 'Bilgisayar Adı',
      version: 'RustDesk Sürümü',
      memory: 'Bellek (RAM)',
      os: 'İşletim Sistemi',
      rustdesk_id: 'RustDesk ID'
    },
    audit: {
      username: 'Kullanıcı',
      type: 'İşlem Türü',
      conn_id: 'Bağlantı ID',
      rustdesk_id: 'Hedef RustDesk ID',
      ip: 'IP Adresi',
      session_id: 'Oturum ID',
      uuid: 'UUID',
      created_at: 'Bağlantı Zamanı',
      closed_at: 'Kapanış Zamanı',
      typeLabel: {
        remote_control: 'Uzak Masaüstü',
        file_transfer: 'Dosya Aktarımı',
        tcp_tunnel: 'TCP Tüneli'
      },
      fileTransferTypeLabel: {
        master_controlled: 'Yönetici -> Hedef Cihaz',
        controlled_master: 'Hedef Cihaz -> Yönetici'
      },
      peer_id: 'Eş Cihaz ID',
      path: 'Dosya Yolu'
    },
    mailTemplate: {
      name: 'Şablon Adı',
      type: 'Tür',
      subject: 'Konu',
      contents: 'İçerik',
      created_at: 'Oluşturulma Tarihi',
      typeLabel: {
        loginVerify: 'Giriş Doğrulama',
        registerVerify: 'Kayıt Doğrulama',
        other: 'Diğer'
      }
    },
    mailLog: {
      username: 'Kullanıcı',
      uuid: 'UUID',
      from: 'Gönderen',
      to: 'Alıcı',
      subject: 'Konu',
      contents: 'İçerik',
      status: 'Durum',
      created_at: 'Gönderim Zamanı',
      statusLabel: {
        ok: 'Başarılı',
        err: 'Hata'
      }
    }
  },
  api: {
    CaptchaError: 'Güvenlik kodu hatası',
    UserNotExists: 'Kullanıcı bulunamadı',
    UsernameOrPasswordError: 'Hatalı kullanıcı adı veya parola',
    UserExists: 'Bu kullanıcı adı zaten kullanımda',
    UsernameEmpty: 'Kullanıcı adı boş bırakılamaz',
    PasswordEmpty: 'Parola boş bırakılamaz',
    UserAddSuccess: 'Kullanıcı başarıyla oluşturuldu',
    DataError: 'Veri hatası oluştu',
    UserUpdateSuccess: 'Kullanıcı bilgileri güncellendi',
    UserDeleteSuccess: 'Kullanıcı başarıyla silindi',
    SessionKillSuccess: 'Oturum başarıyla sonlandırıldı',
    MailTemplateNameEmpty: 'Şablon adı boş bırakılamaz',
    MailTemplateSubjectEmpty: 'Konu boş bırakılamaz',
    MailTemplateContentsEmpty: 'İçerik boş bırakılamaz',
    MailTemplateAddSuccess: 'E-posta şablonu oluşturuldu',
    MailTemplateUpdateSuccess: 'E-posta şablonu güncellendi',
    NoEmailAddress: 'E-posta adresi tanımlanmamış',
    VerificationCodeError: 'Doğrulama kodu hatalı',
    UUIDEmpty: 'UUID boş bırakılamaz'
  }
};

export default local;
