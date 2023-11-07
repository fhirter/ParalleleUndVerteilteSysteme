# Übung Design Goals von verteilten Systemen

## Vorbereitung
Wähle eine der folgenden Technologien:
- [CockroachDB](https://www.cockroachlabs.com/): Datenbank
- [Ably](https://ably.com/): Pub-Sub Messaging
- [Zitadel](https://zitadel.com/): Identity Provider

## Fragen
Beantworte folgende Fragen und erarbeite ein Handout.

### Design-Ziele
Untersuche die Features der gewählten Technologie und vergleiche diese mit den folgenden Design-Zielen von verteilten Systemen:

Transparenz:
- Zugriff
- Standort
- Replikation
- Änderung des Standorts
- Gleichzeitiger Zugriff von verschiedenen Usern
- Fehler

Offenheit:
- Interoperabilität mit anderen Systemen
- Erweiterbarkeit

Zuverlässigkeit
- Verfügbarkeit (Availability, Reliability)
- Sicherheit im Fehlerfall
- Wartbarkeit

Sicherheit 
- Vertraulichkeit und Integrität
- Autorisierung und Authentifizierung

Scalability
- Grössenskalierbarkeit
- Geografische Skalierbarkeit
- Administrative Skalierbarkeit
- Verstecken von Netzwerklatenzen
- Partitionierung und Verteilung
- Replikation

(vgl van Steen 2023 Kapitel 1.2)

## Quellen

van Steen 2023
: M. van Steen and A.S. Tanenbaum, Distributed Systems, 4th ed., distributed-systems.net, 2023.
