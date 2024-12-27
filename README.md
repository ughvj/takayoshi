## Migration tool

Atlas

- inspection example

`atlas schema inspect -u "mysql://takamori:password@localhost:3306/takamori" > schema.hcl`

- application example

`atlas schema apply -u "mysql://takamori:password@localhost:3306/takamori" --to file://schema.hcl`
