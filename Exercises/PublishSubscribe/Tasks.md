# Übung Architekturen

Entwickle und implementiere eine Architektur für eine Chat-Applikation.

Halte dein Vorgehen, Überlegungen und Erkenntnisse in einem Dokument fest.
Schicke dieses am Ende der Unterrichtseinheit dem Dozenten.

Organisiert euch selbstständig in der Klasse, wie ihr die Aufgabe Lösen wollt.
Stellt sicher, dass alle verstehen, was gemacht wurde.

## Lernziele

- Die Studierenden kennen verschiedene Architekturstile und deren Anwendung.
- Die Studierenden könne die technischen Vorzüge spezifischer Entwicklungen im Bereich verteilter Computersysteme
analysieren und nachvollziehbar begründen. Sie können Sachverhalte klar erklären, um die Qualität des Feedbacks zu
optimieren.

## Aufgabenstellung

Halte die Erkenntnisse und Designs in einem [Markdown-Dokument](https://www.markdownguide.org/) mit [Mermaid](https://mermaid.js.org/) fest.

### Architektur

Die Applikation soll Chats zwischen einzelnen Usern sowie in Gruppen ermöglichen.

Dazu bietet sich die Publish-Subscribe in Kombination mit einer "Ports & Adaptors" Architektur an.

Entwirf eine entsprechende Architektur für diese Applikation!

Bestimme ebenfalls das Format der ausgetauschten Nachrichten.

### Implementierung

Nutze für die Implementierung dazu den Message-Broker "RabbitMQ" von [CloudAMQP](https://www.cloudamqp.com/) oder setze
auf einem [Raspberry PI einen MQTT Broker](https://www.elektronik-kompendium.de/sites/raspberry-pi/2709041.htm) auf.

Verbinde den einfachen [Beispielclient](client.html) mit dem Message Broker.

Erweitere die Beispielapplikation:
- Anzeige der verfügbaren User
- Erstellen von Gruppenchats
- Persistierung der Nachrichten in einer Datenbank

## Vorgehen

Softwareprojekte sind sehr oft mit Forschungsprojekten vergleichbar, da oftmals die konkrete Ausgestaltung der Lösung
zu Beginn nicht bekannt ist.
Daher sind klassische Vorgehensmodelle kaum brauchbar.
Es hat sich gezeigt, dass ein Vorgehen, bei dem die Lösung einerseits iterativ verfeinert wird und andererseits
inkremental erweitert wird gute Ergebnisse bringt.

Bei der obigen Aufgabenstellung lassen sich verschiedene Teile identifizieren, die bekannt sind und im Unterricht 
behandelt wurden:
- Publish-Subscribe Architektur
- Ports & Adaptors Architektur
- Beispielanwendung

In Gruppen können diese Themen angegangen werden, entsprechende Architekturentwürfe können gezeichnet werden und die
Beispielapplikation kann in Betrieb genommen werden.

Für den Architekturentwurf bietet sich das [C4-Modell](https://c4model.com/) als Orientierung an. Der Entwurf kann dabei
Top-Down (Context->Containers->Components->Code) oder Bottom-Up erfolgen (Code->Components->Containers->Context).
Es kann auch abgewechselt werden zwischen Top-Down und Bottom-Up.