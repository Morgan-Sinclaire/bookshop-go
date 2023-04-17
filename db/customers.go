package db

func CreateCustomer(name, shippingAddr string) (int, error) {
	database := Connect().Db

	database.Exec(`
		INSERT INTO Customers (name, shippingAddress, accountBalance)
		VALUES (?, ?, 5.0);
	`, name, shippingAddr)

	return GetCustomerId(name, shippingAddr)
}

func GetCustomerId(name, shippingAddr string) (int, error) {
	database := Connect().Db

	rows, err := database.Query(`
		SELECT (id) FROM Customers
		WHERE name = ? and shippingAddress = ?;
	`, name, shippingAddr)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	var cid int
	rows.Next()
	err = rows.Scan(&cid)
	if err != nil {
		return 0, err
	}
	return cid, nil
}

func GetCustomerAddress(cid int) (string, error) {
	database := Connect().Db

	rows, err := database.Query(`
		SELECT (shippingAddress)
		FROM Customers
		WHERE id = ?;
	`, cid)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var addr string
	rows.Next()
	err = rows.Scan(&addr)
	if err != nil {
		return "", err
	}

	return addr, nil
}

func UpdateCustomerAddress(cid int, newAddr string) error {
	database := Connect().Db

	_, err := database.Exec(`
		UPDATE Customers
		SET shippingAddress = ?
		WHERE id = ?;
	`, newAddr, cid)

	return err
}

func CustomerBalance(cid int) (float32, error) {
	database := Connect().Db

	rows, err := database.Query(`
		SELECT (accountBalance)
		FROM Customers
		WHERE id = ?;
	`, cid)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var balance float32
	rows.Next()
	err = rows.Scan(&balance)
	if err != nil {
		return 0, err
	}

	return balance, nil
}

func ChargeCustomerForPO(pid int) error {
	return nil
}
