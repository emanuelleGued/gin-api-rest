package models

import "gorm.io/gorm"

type Aluno struct {
	gorm.Model        //com essa linha ele insere o id, quando foi criado, quando foi atualizado e quando foi deletado
	Nome       string `json:"nome"`
	CPF        string `json:"cpf"`
	RG         string `json:"rg"`
}
