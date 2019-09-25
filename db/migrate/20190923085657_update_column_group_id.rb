class UpdateColumnGroupId < ActiveRecord::Migration[5.2]
  def change
    change_column :users, :group_id, :integer, limit: 8
    change_column :retros, :group_id, :integer, limit: 8
    change_column :reviews, :group_id, :integer, limit: 8
  end
end
