package db

import "errors"

// the abbreviation PO will be used in place of PurchaseOrder often for brevity

func CreatePO(book, customer int) (int, error) {
	database := Connect().Db

	database.Exec(`
		INSERT INTO PurchaseOrders (bookId, customerId, shipped)
		VALUES (?, ?, 0);
	`, book, customer)

	return GetPOByContents(book, customer)
}

func GetPOByContents(book, customer int) (int, error) {
	database := Connect().Db

	rows, err := database.Query(`
		SELECT (id) FROM PurchaseOrders
		WHERE bookId = ? and customerId = ?;
	`, book, customer)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	var pid int
	rows.Next()
	err = rows.Scan(&pid)
	if err != nil {
		return 0, err
	}
	return pid, nil
}

func IsPOShipped(pid int) (bool, error) {
	database := Connect().Db

	rows, err := database.Query(`
		SELECT (shipped) FROM Books
		WHERE id = ?;
	`, pid)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	var shipped int
	rows.Next()
	err = rows.Scan(&shipped)
	if err != nil {
		return false, err
	}
	return shipped == 1, nil
}

func ShipPO(pid int) error {
	database := Connect().Db

	if isShipped, err := IsPOShipped(pid); err == nil && isShipped {
		return errors.New("purchase order already shipped")
	} else if err != nil {
		return err
	}

	err := ChargeCustomerForPO(pid)
	if err != nil {
		return err
	}

	_, err = database.Exec(`
		UPDATE PurchaseOrders
		SET shipped = 1
		WHERE id = ?;
	`, pid)

	return err
}
