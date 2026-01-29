# WhatsX - Multi-Instance WhatsApp Wrapper

[Indonesian](https://github.com/Dlcuy22/WhatsX/blob/main/README.id.md)

WhatsX is a lightweight wrapper for [web.whatsapp.com](https://web.whatsapp.com) built using Wails. Its core feature is **multi-instance** support, allowing you to run multiple WhatsApp accounts simultaneously with isolated profiles.

## Features

- **Multi-Instance**: Run multiple WhatsApp accounts at once (Personal, Business, etc.).
- **Data Isolation**: Each profile has its own storage, ensuring sessions do not mix.
- **Portable**: Data is stored locally within the application folder (`data/`), keeping your system clean.
- **Lightweight**: Utilizes the system's native WebView (WebView2 on Windows).

## Tech Stack

- **Backend**: [Wails](https://wails.io/) (Golang) - Handles window management and the multi-instance system.
- **Frontend**: React.js - Handles the loading interface and WebView initialization.

### Running the Application

By default, running the application without arguments will load the `default` profile.

```bash
WhatsX
```

To run a specific profile:

```bash
WhatsX --profile <profile_name>
```

Examples:

```bash
WhatsX --profile business
WhatsX --profile gaming
```

## Configuration

Profiles are configured via the `WhatsX.config.json` file in the application's root directory. You can add new instances simply by editing this file.

**Example `WhatsX.config.json` Structure:**

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

## Data Structure

This application is designed to be clean and portable.

- All data is stored inside the `data/` folder relative to the executable.
- Each profile has its own sub-folder (e.g., `data/personal`, `data/business`).
- Inside these folders lies the WebView data (cookies, local storage, WhatsApp session cache).

Structure example:

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

- [ ] Native System Notification Integration
- [ ] Standalone GUI Launcher for Instance Management

## License

This software is licensed under **GPLv3**.
