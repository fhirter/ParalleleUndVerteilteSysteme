# Übung Architekturen

Entwickle und implementiere eine Architektur für eine Chat-Applikation.

Halte die Erkenntnisse und Designs in einem Dokument fest. Schicke dieses am Ende der Unterrichtseinheit dem Dozenten.

## Lernziele

- Die Studierenden kennen verschiedene Architekturstile und deren Anwendung.
- Die Studierenden können Sachverhalte klar erklären, um die Qualität des Feedbacks zu optimieren.

## Aufgabenstellung

Die Applikation soll Chats zwischen einzelnen Usern sowie in Gruppen ermöglichen.

Dazu bietet sich die Publish-Subscribe in Kombination mit einer "Ports & Adaptors" Architektur an.

Entwirf eine entsprechende Architektur für diese Applikation!

Bestimme ebenfalls das Format der ausgetauschten Nachrichten.

### Implementierung

Nutze für die Implementierung einen der folgenden Message-Broker

- "RabbitMQ" von [CloudAMQP](https://www.cloudamqp.com/)
- ["Eclipse Mosquitto"](https://mosquitto.org/)
- [NATS](https://nats.io/)

Verbinde den einfachen [Beispielclient](client.html) mit dem Message-Broker.

Erweitere die Beispielapplikation:

- Anzeige der aktiven User
- Erstellen von Gruppenchats
- Persistierung der Nachrichten in einer Datenbank

## Vorgehen

Softwareprojekte sind sehr oft mit Forschungsprojekten vergleichbar, da oftmals die konkrete Ausgestaltung der Lösung zu
Beginn nicht bekannt ist.

Es hat sich gezeigt, dass ein Vorgehen, bei dem die Lösung iterativ verfeinert und inkrementell erweitert wird, gute
Ergebnisse bringt.

Bei der obigen Aufgabenstellung lassen sich verschiedene Teile identifizieren:

- Publish-Subscribe Architektur
- Ports & Adaptors Architektur
- Beispielanwendung
- Erweiterungen

In Gruppen können diese Themen angegangen werden, entsprechende Architekturentwürfe können gezeichnet werden und die
Beispielapplikation kann in Betrieb genommen werden.

Für den Architekturentwurf bietet sich das [C4-Modell](https://c4model.com/) an.
Ein Top-Down Vorgehen (Context->Containers->Components->Code) ist meistens sinnvoll.