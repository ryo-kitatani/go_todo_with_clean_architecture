data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./internal/domain/models",
    "--dialect", "mysql",
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  dev = "mysql://root:password@localhost:3306/todos"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
