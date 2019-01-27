class ConvertToUtfmb4 < ActiveRecord::Migration[5.2]
  def change
    # for each table that will store unicode execute:
    execute "ALTER TABLE retros CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"
    # for each string/text column with unicode content execute:
    execute "ALTER TABLE retros CHANGE message message VARCHAR(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"
  end
end
