class CreateGroups < ActiveRecord::Migration[5.2]
  def change
    create_table :groups, id: false, primary_key: :id do |t|
      t.primary_key :id, :unsigned_integer, auto_increment: true
      t.integer :chat_id, null: false
      t.string :name, null: false
    end
  end
end
