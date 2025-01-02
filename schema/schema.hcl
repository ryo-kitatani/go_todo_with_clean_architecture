schema "main" {}

table "todos" {
 schema = schema.main
 column "id" {
   type = bigint
   unsigned = true
   auto_increment = true
   null = false
 }
 
 index "pk_todos" {
   on = table.todos.id
   unique = true
 }
 
 column "title" {
   type = varchar(255)
   null = false
 }
 column "completed" {
   type = bool
   default = false
 }
 column "created_at" {
   type = timestamp
   null = false
 }
 column "updated_at" {
   type = timestamp
 }
 column "deleted_at" {
   type = timestamp
 }
}