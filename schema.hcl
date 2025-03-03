table "students" {
  schema = schema.aps
  column "id" {
    null           = false
    type           = int
    unsigned       = true
    auto_increment = true
  }
  column "name" {
    null           = false
    type           = varchar(1023)
  }
  column "nyutaikun_user_id" {
    null           = false
    type           = int
    unsigned       = true
  }

  primary_key {
    columns = [column.id]
  }
}

table "inspected_data" {
  schema = schema.aps
  column "id" {
    null           = false
    type           = int
    unsigned       = true
    auto_increment = true
  }
  column "students_id" {
    null           = false
    type           = int
    unsigned       = true
  }
  column "mental_data" {
    null           = false
    type           = int
    unsigned       = true
  }
  column "physical_data" {
    null           = false
    type           = int
    unsigned       = true
  }
  column "created_at" {
    null = false
    type = date
  }


  primary_key {
    columns = [column.id]
  }
}

table "t_inspected_history" {
  schema = schema.aps
  column "id" {
    null           = false
    type           = int
    unsigned       = true
    auto_increment = true
  }
  column "nyutaikun_user_id" {
    null           = false
    type           = int
    unsigned       = true
  }

  primary_key {
    columns = [column.id]
  }
}

schema "aps" {
  charset = "utf8mb4"
  collate = "utf8mb4_0900_ai_ci"
}
