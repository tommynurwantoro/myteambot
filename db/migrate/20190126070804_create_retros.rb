class CreateRetros < ActiveRecord::Migration[5.2]
  def change
    create_table :retros, id: false, primary_key: :id do |t|
      t.primary_key :id, :unsigned_integer, auto_increment: true
      t.string :username, null: false
      t.string :type, null: false #glad,sad,mad
      t.string :message, null: false
      
      t.timestamps null: false
    end
  end
end
