package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	// Connexion au serveur 172.16.0.2 sur le port 8000 en TCP
	conn, err := net.Dial("tcp", "172.16.0.2:8000")
	if err != nil {
		fmt.Println("Erreur de connexion :", err) // Renvoie un message d'erreur si le programme n'arrive pas à se connecter
		return
	}
	defer conn.Close() // S'il y a une erreur, ferme la connexion

	// Goroutine pour écouter les réponses du serveur en continu
	go func() {
		reader := bufio.NewReader(conn)
		for {
			message, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("\nConnexion perdue avec le serveur")
				os.Exit(1)
			}
			fmt.Printf("\nRéponse du serveur : %s", message)
			fmt.Print("Message à envoyer : ")
		}
	}()

	// Boucle principale pour envoyer les messages
	inputReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("\nMessage à envoyer : ")
		text, _ := inputReader.ReadString('\n')

		// Mesurer la latence avant l'envoi
		startTime := time.Now()

		// Envoie du message au serveur qui écoute sur le port 8000
		_, err := conn.Write([]byte(text))
		if err != nil {
			fmt.Println("Erreur d'envoi :", err)
			break
		}

		// Attendre un court instant pour éviter un bug d'affichage
		time.Sleep(100 * time.Millisecond)

		// Afficher la latence immédiatement après l'envoi
		latency := time.Since(startTime)
		fmt.Printf("Latence : %v\n", latency)
	}
}
