package rCassandra

func (r *repository) CloseConnection() {
	if r.session != nil {
		r.session.Close()
	}
}
