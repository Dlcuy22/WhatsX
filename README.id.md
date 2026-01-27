# WhatsX - Multi-Instance WhatsApp Wrapper

WhatsX adalah aplikasi wrapper ringan untuk [web.whatsapp.com](https://web.whatsapp.com) yang dibangun menggunakan Wails. Fitur utamanya adalah kemampuan **multi-instance**, memungkinkan Anda menjalankan beberapa akun WhatsApp secara bersamaan dengan profil yang terisolasi.

## Fitur

- **Multi-Instance**: Jalankan banyak akun WhatsApp sekaligus (Personal, Bisnis, dll).
- **Isolasi Data**: Setiap profil memiliki penyimpanan data sendiri, tidak tercampur.
- **Portable**: Data disimpan secara lokal di dalam folder aplikasi (`data/`), tidak menyebar di sistem operasi.
- **Ringan**: Menggunakan WebView asli sistem (WebView2 di Windows).

## Stack Teknologi

- **Backend**: [Wails](https://wails.io/) (Golang) - Menangani window management dan sistem multi-instance.
- **Frontend**: React.js - Menangani antarmuka loading dan inisialisasi WebView.

### Menjalankan Aplikasi

Secara default, jika dijalankan tanpa argumen, WhatsX akan memuat profil `default`.

```bash
WhatsX
```

Untuk menjalankan profil tertentu::

```bash
WhatsX --profile <nama_profil>
```

Contoh:

```bash
WhatsX --profile business
WhatsX --profile gaming
```

## Konfigurasi

Profil dikonfigurasi melalui file `WhatsX.config.json` di root folder aplikasi. Anda dapat menambahkan instance baru cukup dengan mengedit file ini.

**Contoh Struktur `WhatsX.config.json`:**

```json
{
  "profiles": {
    "default": {
      "name": "Personal",
      "data_path": "data/personal"
    },
    "business": {
      "name": "Business Account",
      "data_path": "data/business"
    },
    "gaming": {
      "name": "Gaming Community",
      "data_path": "data/gaming"
    }
  }
}
```

## Struktur Data

Aplikasi ini didesain agar rapi dan tidak mengotori sistem file Anda.

- Semua data disimpan di dalam folder `data/` relatif terhadap lokasi executable.
- Setiap profil memiliki sub-folder sendiri (misal: `data/personal`, `data/business`).
- Di dalam folder tersebut terdapat data WebView (cookies, local storage, cache session WhatsApp).

Contoh struktur:

```
WhatsX/
├── WhatsX.exe
├── WhatsX.config.json
└── data/
    ├── personal/
    │   └── EBWebView/ ...
    └── business/
        └── EBWebView/ ...
```

## Roadmap

- [ ] Integrasi Notifikasi Sistem (Native System Notification)
- [ ] Launcher GUI Berdiri Sendiri untuk Manajemen Instance

## Lisensi

Software ini dilisensikan di bawah **GPLv3**.
