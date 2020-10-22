package main

import (
	"fmt"
	"log"
)

type account struct {
	name string
}

func newAccount(accountName string) *account{
	return &account{
		name: accountName,
	}
}

func (a *account) checkAccount(accountName string) error {
	if a.name != accountName{
		return fmt.Errorf("Account name is incorrect")
	}
	fmt.Println("Account Verified")
	return nil
}

type securityCode struct {
	code int
}

func newSecurityCode(code int) *securityCode{
	return &securityCode{
		code: code,
	}
}

func (s *securityCode) checkSecurityCode(incomingCode int) error {
	if s.code != incomingCode{
		return fmt.Errorf("Security code is incorrect")
	}
	fmt.Println("SecurityCode Verified")
	return nil
}

type wallet struct {
	balance int
}

func newWallet() *wallet{
	return &wallet{
		balance: 0,
	}
}

func (w *wallet) creditBalance(amount int)  {
	w.balance += amount
	fmt.Println("Wallet balance added successfully")
}

func (w *wallet) debitBalance(amount int) error {
	if w.balance < amount{
		return fmt.Errorf("Balance is not sufficient")
	}
	fmt.Println("Wallet balance is Sufficient")
	w.balance = w.balance - amount
	return nil
}


type ledger struct {
}

func (s *ledger) makeEmpty(accountID, txnType string, amount int)  {
	fmt.Printf("Make ledger entry for accountId %s with txnType %s for amount %d\n", accountID, txnType, amount)
	return
}

type notification struct {
}

func (n *notification) sendWalletCreditNotification() {
	fmt.Println("Sending wallet credit notification")
}

func (n *notification) sendWalletDebitNotification() {
	fmt.Println("Sending wallet debit notification")
}

type walletFacade struct {
	account *account
	wallet *wallet
	securityCode *securityCode
	notification *notification
	ledger *ledger
}

func newWalletFacade(accountID string, code int) *walletFacade{
	fmt.Println("Starting create account")
	w := &walletFacade{
		account: newAccount(accountID),
		securityCode: newSecurityCode(code),
		wallet: newWallet(),
		notification: &notification{},
		ledger: &ledger{},
	}
	fmt.Println("Account created")
	return w
}

func (w *walletFacade) addMonetToWallet(accountID string, securityCode , amount int) error {
	fmt.Println("Starting add money to wallet")
	err := w.account.checkAccount(accountID)
	if err != nil{
		return err
	}
	err = w.securityCode.checkSecurityCode(securityCode)
	if err != nil{
		return err
	}
	w.wallet.creditBalance(amount)
	w.notification.sendWalletCreditNotification()
	w.ledger.makeEmpty(accountID, "credit,", amount)
	return nil
}

func (w *walletFacade) deductMonetFromWallet(accountID string, securityCode , amount int) error {
	fmt.Println("Starting debit money to wallet")
	err := w.account.checkAccount(accountID)
	if err != nil{
		return err
	}
	err = w.securityCode.checkSecurityCode(securityCode)
	if err != nil{
		return err
	}
	w.wallet.debitBalance(amount)
	w.notification.sendWalletDebitNotification()
	w.ledger.makeEmpty(accountID, "debit,", amount)
	return nil
}

func main()  {
	wFacade := newWalletFacade("abc", 1234)
	err := wFacade.addMonetToWallet("abc", 1234, 15)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
	fmt.Println()
	err = wFacade.deductMonetFromWallet("abc", 1234, 5)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

}