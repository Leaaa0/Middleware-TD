# Service Scheduler

Service responsable de la récupération périodique des événements d'emploi du temps (ICal) et de leur publication dans NATS JetStream.

## Prérequis

- **NATS Server** avec JetStream activé.
- **Service Config** doit être lancé pour fournir les IDs d'agendas (`http://localhost:8091`).

## Démarrage

### 1. Lancer NATS Server
Dans un terminal séparé :
```bash
nats-server -js
```

### 2. Écouter les événements (Debug)
Pour vérifier que les événements sont bien publiés :
```bash
nats sub "EVENTS.>"
```

### 3. Lancer le Scheduler
Mettre à jour les dépendances :
```bash
go mod tidy
```

Lancer le service :
```bash
go run cmd/main.go
```

Ou compiler et lancer :
```bash
go build -o scheduler cmd/main.go
./scheduler
```

## Fonctionnement

Le Scheduler effectue les actions suivantes :
1. Se connecte à NATS JetStream.
2. Récupère la liste des IDs d'agenda via l'API Config (`/calendars`).
3. Télécharge le fichier ICal correspondant depuis `edt.uca.fr`.
4. Parse le fichier ICal pour extraire les événements.
5. Publie chaque événement sur le sujet `EVENTS.new` dans NATS.

La synchronisation est effectuée :
- Immédiatement au démarrage.
- Puis toutes les 10 minutes.

## Configuration

- **Port API Config** : 8091 (codé en dur pour le moment).
- **Sujet NATS** : `EVENTS.new`.
- **Intervalle** : 10 minutes.
