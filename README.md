# go-bank
API REST go pour gestion de compte basique

## Installation
go get -u github.com/tipounet/go-bank

## Configuration
La connexion à la base de données ainsi que le port d'écoute HTTP sont configurables dans le fichier application.yaml (le context http n'est pas encore pris en compte).

## Description des service

### Bank
``` json
{
   "id":5,
   "name":"ma bank"
}
```
#### Services disponibless

Path | Méthode | Description
-----|---------|-------------
/bank | GET | retournes la liste de toutes les banques
bank/{id} | GET | retourne la banque correspondant à l'id
bank/name/{name} |GET | Retourne les banques  dont le nom correspond (name like %{name%})
/bank | POST | Création du banque avec le JSON dans le corps de la requête
/bank | PUT | Mise à jour d'une banque avec le JSON dans le corps de la requête
/bank/{id} | DELETE | Suppression d'une banque à partir de son ID

### Account : un compte bancaire
``` json
{
   "id":1,
   "number":"132456789azerty",
   "user":{
      "id":76,
      "Nom":"user name",
      "prenom":"user first name",
      "pseudo":"pseudouser",
      "email":"monem@il.com",
      "pwd":"passwd",
      "pwdbit":null
   },
   "bank":{
      "id":5,
      "name":"ma bank"
   }
}
```
#### Services disponibles

Path | Méthode | Description
-----|---------|-------------
/account | GET | retournes la liste de tous les comptes
/account/{id} | GET | retourne le compte correspondant à l'id
/account | POST | Création d'un compte avec le JSON dans le corps de la requête
/account | PUT | Mise à jour d'un compte avec le JSON dans le corps de la requête
/account/{id} | DELETE | Suppression d'un compte à partir de son ID
/account/number/{id} | DELETE | Suppression d'un compte à partir de son numéro

### User : un utilisateur de l'application
``` json
{
   "id":76,
   "Nom":"User name",
   "prenom":"user first name",
   "pseudo":"psuedouser",
   "email":"monem@il.com",
   "pwd":"passwd",
   "pwdbit":""
}
```
#### Services disponibles

Path | Méthode | Description
-----|---------|-------------
/user | GET | retournes la liste de tous les utilisateurs
/user/{id} | GET | retournes l'utilisateur correspondant à l'id
/user | POST | Création d'un utilisateur avec le JSON dans le corps de la requête
/user | PUT | Mise à jour d'un utilisateur avec le JSON dans le corps de la requête
/user/{id} | DELETE | Suppression d'un utilisateur à partir de son ID

### Transaction : une transaction sur un compte c'est crédit ou un débit, montant, un id etc.
``` json
{
   "id":1,
   "description":"ma première transaction",
   "Posteddate":"2016-07-27T00:00:00Z",
   "userdate":"2016-07-27T00:00:00Z",
   "fiid":"xyzdfrezsastyh",
   "amount":42.42,
   "account":{
      "id":1,
      "number":"12345789x987564",
      "user":{
         "id":76,
         "Nom":"userName",
         "prenom":"user first name",
         "pseudo":"pseudouser",
         "email":"monem@il.com",
         "pwd":"passwd",
         "pwdbit":null
      },
      "bank":{
         "id":5,
         "name":"ma banque"
      }
   },
   "type":{
      "id":1,
      "name":"Crédit"
   }
}
```
#### Services disponibles

Path | Méthode | Description
-----|---------|-------------
/transaction | GET | retournes la liste de toutes les transactions
/transaction/{id} | GET | retourne la transaction correspondant à l'id
/transaction | POST | Création d'une transaction avec le JSON dans le corps de la requête
/transaction | PUT | Mise à jour d'une transaction avec le JSON dans le corps de la requête
/transaction/{id} | DELETE | Suppression d'une transaction à partir de son ID

### TransactionType : un type de transaction : crédit / débit
``` json
{
   "id":1,
   "name":"Crédit"
}
```
#### Services disponibles

Path | Méthode | Description
-----|---------|-------------
/transactionType | GET | retournes la liste de tous les types de transaction
/transactionType/{id} | GET | retourne le type de transaction correspondant à l'id
/transactionType | POST | Création d'un type de transaction avec le JSON dans le corps de la requête
/transactionType | PUT | Mise à jour d'un type de transaction avec le JSON dans le corps de la requête
/transactionType/{id} | DELETE | Suppression d'un type de transaction à partir de son ID

# TODO

* i18n
* webapp dans le projet
* page index avec description du projet (swagger ?)
* authentification avec JWT entre font et back ?
* Gestion des services avec et sans authentifications
* Tests unitaires
* Meilleur gestion des erreurs go renvoyées (code, format etc.)
