package main // Déclare le package principal pour le programme

import (
    "fmt" // Le package fmt permet d'afficher des messages à la console.
    "log" // Le package log est utilisé pour enregistrer des messages de log, comme des erreurs.
    "example.com/greetings" // C'est un import pour un package personnalisé que tu as créé (greetings). 
                              // Il contient probablement une fonction `Hello`.
)

func main() { // Fonction principale où l'exécution du programme commence.
    // Initialiser le préfixe et les options de format du package log.
    log.SetPrefix("greetings: ") // Le préfixe "greetings: " sera ajouté au début de chaque message de log.
    log.SetFlags(0) // Supprime les options de formatage par défaut dans les messages de log (comme la date/heure).

    // Appel de la fonction `Hello` du package `greetings` pour obtenir un message de salutation.
    // L'argument "" (chaîne vide) est passé, donc on s'attend à ce que `Hello` gère cela (probablement avec un message par défaut ou une erreur).
    // message, err := greetings.Hello("Wifak") // On attend un message de type string et une erreur possible.

		names := []string {"Wifak","Fanny","Bassirath"}
		messages , err := greetings.Hellos(names)
    if err != nil { // Si une erreur se produit (par exemple, si le nom est vide ou non valide dans la fonction `Hello`),
        log.Fatal(err) // On enregistre l'erreur et on termine le programme immédiatement.
    }

    // Si aucune erreur ne se produit, on affiche le message retourné par la fonction `Hello`.
    fmt.Println(messages) // Affiche le message dans la console.
}
