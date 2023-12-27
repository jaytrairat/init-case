
# init-case
A command line for creating a case folder and evidence folders.

The script will create the parent folder and sub-folder for evidence, and then create the following sub-folders inside the evidence folder:

- Documents
- Images&Extractions
- Pictures
- Reports
- Request
- Videos

## Installation
```bash
go install github.com/jaytrairat/init-case@V0.0.3
```

## Usage
```bash
init-case -c 1 -e 1 -e 2 -y 2566
```

<!-- ## Example -->
<!-- ![Example](https://raw.githubusercontent.com/jaytrairat/init-case/main/assets/demo.gif) -->

## Result
```bash
    F-2566-001/
    ├─ Documents/
    │   ├─ F-2566-001-EV01/
    │   └─ F-2566-001-EV02/
    ├─ Images&Extractions/
    │   ├─ F-2566-001-EV01/
    │   └─ F-2566-001-EV02/
    ├─ Pictures/
    │   ├─ F-2566-001-EV01/
    │   └─ F-2566-001-EV02/
    ├─ Reports/
    │   ├─ F-2566-001-EV01/
    │   └─ F-2566-001-EV02/
    ├─ Request/
    │   ├─ F-2566-001-EV01/
    │   └─ F-2566-001-EV02/
    └─ Videos/
        ├─ F-2566-001-EV01/
        └─ F-2566-001-EV02/
```
