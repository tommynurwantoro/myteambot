class RefactorColumnTableReview < ActiveRecord::Migration[5.2]
  def change
    rename_column :reviews, :is_done, :is_reviewed
    add_column :reviews, :is_tested, :boolean, null: false, default: false
    change_column :reviews, :group_id, :integer, null: false, limit: 8
  end
end
