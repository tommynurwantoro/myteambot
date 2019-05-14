class UpdateReviews < ActiveRecord::Migration[5.2]
  def change
    add_column :reviews, :title, :string, null: false
    add_column :reviews, :users, :string, null: false
  end
end
