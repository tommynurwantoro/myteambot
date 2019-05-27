package method

/**
start - Tentang bot ini
help - Nampilin semua perintah yang ada
halo - Cuma buat nyapa aja
init_group - Initialize group ke dalam bot
add_user - Menambah user agar dapat menggunakan bot
retro - Bantuan untuk perintah retrospective
result_retro - {dd-mm-yyyy} Dapetin hasil retrospective, jangan lupa kasih tanggalnya ya
titip_review - {title#url#telegram-users} Titip review PR
antrian_review - Nampilin semua antrian PR yang belum direview
sudah_direview - {urutan} Ngubah antrian review untuk yang sudah direview
sudah_direview_semua - {urutan} Ngubah antrian review untuk yang sudah direview untuk semua user
tambah_user_review - {urutan#telegram-users} Nambahin user ke antrian review
antrian_qa - Nampilin semua antrian PR yang belum dites
sudah_dites - {urutan} Ngubah antrian QA untuk yang sudah dites
**/

// Command _
type Command struct {
	Name        string
	Description string
}

func GetCommand() *Command {
	return &Command{}
}

func (c *Command) All() []Command {
	// AllCommands List all commands
	return []Command{
		c.Start(),
		c.Help(),
		c.Halo(),
		c.InitGroup(),
		c.AddUser(),
		c.Retro(),
		c.Glad(),
		c.Sad(),
		c.Mad(),
		c.ResultRetro(),
		c.TitipReview(),
		c.AntrianReview(),
		c.SudahDireview(),
		c.SudahDireviewSemua(),
		c.TambahUserReview(),
		c.AntrianQA(),
		c.SudahDites(),
	}
}

func (c *Command) Start() Command {
	return Command{"/start", "Tentang bot ini"}
}

func (c *Command) Help() Command {
	return Command{"/help", "Nampilin semua perintah yang ada"}
}

func (c *Command) Halo() Command {
	return Command{"/halo", "Cuma buat nyapa aja"}
}

func (c *Command) Retro() Command {
	return Command{"/retro", "Bantuan untuk perintah retrospective"}
}

func (c *Command) Glad() Command {
	return Command{"/glad", "Pesan glad untuk retro"}
}

func (c *Command) Sad() Command {
	return Command{"/sad", "Pesan sad untuk retro"}
}

func (c *Command) Mad() Command {
	return Command{"/mad", "Pesan mad untuk retro"}
}

func (c *Command) ResultRetro() Command {
	return Command{"/result_retro", "{dd-mm-yyyy} Dapetin hasil retrospective, jangan lupa kasih tanggalnya ya"}
}

func (c *Command) TitipReview() Command {
	return Command{"/titip_review", "{title#url#telegram-users} Titip review PR"}
}

func (c *Command) AntrianReview() Command {
	return Command{"/antrian_review", "Nampilin semua antrian PR yang belum direview"}
}

func (c *Command) SudahDireview() Command {
	return Command{"/sudah_direview", "{urutan} Ngubah antrian review untuk yang sudah direview"}
}

func (c *Command) SudahDireviewSemua() Command {
	return Command{"/sudah_direview_semua", "{urutan} Ngubah antrian review untuk yang sudah direview untuk semua user"}
}

func (c *Command) TambahUserReview() Command {
	return Command{"/tambah_user_review", "{urutan#telegram-users} Nambahin user ke antrian review"}
}

func (c *Command) AntrianQA() Command {
	return Command{"/antrian_qa", "Nampilin semua antrian PR yang belum dites"}
}

func (c *Command) SudahDites() Command {
	return Command{"/sudah_dites", "{urutan} Ngubah antrian QA untuk yang sudah dites"}
}

func (c *Command) InitGroup() Command {
	return Command{"/init_group", "Initialize group ke dalam bot"}
}

func (c *Command) AddUser() Command {
	return Command{"/add_user", "Menambah user agar dapat menggunakan bot"}
}
