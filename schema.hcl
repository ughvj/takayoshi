table "questions" {
  schema = schema.takamori
  column "id" {
    null           = false
    type           = int
    unsigned       = true
    auto_increment = true
  }
  column "statement" {
    null           = false
    type           = varchar(1023)
  }
  column "category" {
    null           = false
    type           = varchar(30)
  }

  primary_key {
    columns = [column.id]
  }
}

table "question_options" {
  schema = schema.takamori
  column "id" {
    null           = false
    type           = int
    unsigned       = true
    auto_increment = true
  }
  column "correct_choice" {
    null           = true
    type           = boolean
  }
  column "correct_order" {
    null           = true
    type           = int
    unsigned       = true
  }
  column "question_id" {
    null           = false
    type           = int
    unsigned       = true
  }
  column "genkun_id" {
    null           = false
    type           = int
    unsigned       = true
  }

  primary_key {
    columns = [column.id]
  }
}

table "genkuns" {
  schema = schema.takamori
  column "id" {
    null           = false
    type           = int
    unsigned       = true
    auto_increment = true
  }
  column "name_kanji" {
    null           = false
    type           = varchar(30)
  }
  column "name_yomi_hiragana" {
    null           = false
    type           = varchar(255)
  }
  column "src" {
    null           = false
    type           = varchar(255)
  }

  primary_key {
    columns = [column.id]
  }
}
schema "takamori" {
  charset = "utf8mb4"
  collate = "utf8mb4_0900_ai_ci"
}
