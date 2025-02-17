package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	// Écoute sur le port 8000
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("Erreur d'écoute :", err)
		return
	}
	defer listener.Close()
	fmt.Println("Serveur en écoute sur le port 8000")

	for {
		// Accepte une connexion
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Erreur d'acceptation :", err)
			continue
		}
		fmt.Println("Nouvelle connexion :", conn.RemoteAddr())

		// Gère la connexion dans une goroutine (parallèle)
		go handleConnection(conn)
	}
}

// 🔹 Fonction pour gérer chaque client connecté
func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		// Lire le message du client
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Connexion fermée par le client")
			return
		}

		fmt.Println("Message reçu :", message)

		// Envoyer une réponse de confirmation
		response := "\nMessage bien reçu"
		conn.Write([]byte(response))
	}
}
