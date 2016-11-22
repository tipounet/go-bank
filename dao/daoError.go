package dao

// DAOerror : Message d'erreur lors d'une requête
type DAOerror struct {
	Code    int64
	Message string
}

// Implemente l'interface error et du coup on peur retourner une DAOerror comme error
func (e *DAOerror) Error() string {
	return e.Message
}
