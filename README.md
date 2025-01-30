# Go Download Env Tool

## Software Flow

![software-arch](./public/software_arch.png)

## Infra Architecture

![infra-arch](./public/infra-arch.png)

## RBAC

| Role     | Admin | Developer | 개발 완료 |
| -------- | ----- | --------- | --------- |
| setting  | O     | X         | O         |
| preview  | O     | X         | O         |
| insert   | O     | X         | X         |
| select   | O     | O         | X         |
| update   | O     | X         | X         |
| delete   | O     | X         | X         |
| download | O     | O         | X         |
