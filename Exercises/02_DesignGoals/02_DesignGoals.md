# Übung Design Goals von verteilten Systemen

## Vorbereitung
Wähle eine der folgenden Technologien:
- CockroachDB
- Matrix (Kommunikationsprotokoll)
- Mastodon
- Ably
- Neon (Datenbank)

## Fragen
Beantworte folgende Fragen und erarbeite ein Handout.

### Distributed vs Decentralized
Was ist der Grund für die Verteilung auf verschiedene Systeme? 
Werden bestehende Systeme vernetzt oder ein bestehendes System verteilt?
(vgl van Steen 2023 Kapitel 1.1)

### Design Goals
Untersuche die Technologie auf folgende Eigenschaften:
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

(vgl van Steen 2023 Kapitel 1.2)

## Quellen
van Steen 2023
: M. van Steen and A.S. Tanenbaum, Distributed Systems, 4th ed., distributed-systems.net, 2023.
