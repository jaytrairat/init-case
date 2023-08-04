
# init-case
A command line for creating a case folder and evidence folders.


The `--case` argument is optional and will default to `YYYYMMDDHHmmss_NEWCASE` if not provided. 

The `--evidence` argument is also optional and will default to `EVIDENCE` if not provided.

The script will create the parent folder and sub-folder for evidence, and then create the following sub-folders inside the evidence folder:

- Documents
- Extracts
- Pictures
- Videos
- Reports
- CaseFiles

## Installation
```bash
go install github.com/jaytrairat/init-case@latest
```

## Usage
```bash
init-case -c "F01-66" -e "EV01" -e "EV02"
```

## Example
![Example](https://raw.githubusercontent.com/jaytrairat/init-case/main/assets/demo.gif)

## Result
```bash
    F01-66_NEWCASE/
    ├─ Documents/
    │   ├─ EV01/
    │   └─ EV02/
    ├─ Extracts/
    │   ├─ EV01/
    │   └─ EV02/
    ├─ Pictures/
    │   ├─ EV01/
    │   └─ EV02/
    ├─ Videos/
    │   ├─ EV01/
    │   └─ EV02/
    ├─ Reports/
    │   ├─ EV01/
    │   └─ EV02/
    └─ CaseFiles/
        ├─ EV01/
        └─ EV02/
```
