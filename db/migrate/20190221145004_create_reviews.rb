class CreateReviews < ActiveRecord::Migration[5.2]
  def change
    create_table :reviews, id: false, primary_key: :id do |t|
      t.primary_key :id, :unsigned_integer, auto_increment: true
      t.integer :group_id, null: false, limit: 8
      t.string :url, null: false
      t.string :title, null: false
      t.boolean :is_reviewed, null: false
      t.boolean :is_tested, null: false
      t.string :users, null: false

      t.timestamps null: false
    end
  end
end
