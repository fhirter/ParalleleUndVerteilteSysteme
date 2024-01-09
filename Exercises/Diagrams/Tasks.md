# Übung Diagramme

## Lernziele

- Die Studierenden kennen verschiedene Architekturstile und deren Anwendung.
- Die Studierenden können Sachverhalte klar erklären, um die Qualität des Feedbacks zu optimieren.

## Aufgabenstellung

Erarbeite in Zweiergruppen ein Diagramm einer der folgend beschriebenen Systemarchitekturen.
Ziel ist, die Architektur ohne zusätzlichen Erläuterungen verständlich zu erklären.

- Wähle einen geeigneten Diagrammtyp.
- Wähle eine geeignete Form (Diagram as Code, Zeichenprogramm, Papier).
- Erstelle das Diagramm.

Die Diagramme werden anschliessend einzeln betrachtet.
Die Zeit ist dabei auf 2 Minuten pro Diagramm begrenzt.
Halte dir dabei fest, was gut und was weniger gut verständlich ist.

### Systemarchitekturen

#### Online-Shop mit Microservices

Ein Online-Shop besteht aus mehreren Microservices. Der Inventar-Service verwaltet das Inventar der Artikel, die im
Shop angeboten werden. Die Mitarbeitenden im Lager arbeiten mit dem Inventar-Service, zur Lagerverwaltung.

Der Bestell-Service verarbeitet die Bestellung der Kunden. Die angebotenen Artikel stammen aus dem Inventar-Service.
Für die Abwicklung der Bezahlung wird auf einen externen Zahlungsdienstleister zugegriffen.

Der Liefer-Service verarbeitet die Bestellung nach erfolgreicher Bezahlung. Er greift auf den Kunden-Service zu, um die
Lieferadresse zu erhalten. Der Bestell-Service wird über den erfolgreichen Versand informiert.

#### Online-Shop mit Client-Server-Architektur

Ein E-Commerce-Website, die auf einem zentralen Server gehostet wird.
Kunden nutzen verschiedene Geräte (Smartphones, Laptops) als Clients, um auf die Website zuzugreifen und Bestellungen zu
tätigen.
Der Server verarbeitet Anfragen, verwaltet Nutzerkonten und führt Transaktionen aus.
Ein Datenbankserver speichert Informationen über Produkte, Nutzer und Bestellungen.
Ein optionaler Lastverteiler kann eingesetzt werden, um den Verkehr auf mehrere Server zu verteilen.

#### Streaming-Dienst mit Microservices-Architektur

Ein Video-Streaming-Dienst wie Netflix, der eine Microservices-Architektur verwendet.
Verschiedene Microservices sind verantwortlich für Benutzerauthentifizierung, Video-Streaming, Empfehlungsalgorithmen
und Nutzerdatenverwaltung.
Jeder Microservice ist unabhängig und kommuniziert über ein Netzwerk mit den anderen Services.
Jeder Service verwendet eine eigene Datenbank.
Der Video-Streaming-Service ist global verteilt, um den globalen Traffic zu minimieren und niedrige Ladezeiten zu
bieten.

#### Intelligentes Verkehrssystem mit Edge-Computing

Ein System zur Verkehrsüberwachung und -steuerung in einer Stadt.
Sensoren an Ampeln und auf Straßen sammeln Echtzeitdaten über den Verkehr.
Die Daten werden teilweise direkt vor Ort (Edge-Geräte) verarbeitet, um schnelle Reaktionszeiten zu
ermöglichen.
Zentrale Server sammeln aggregierte Daten für weiterführende Analysen und langfristige Planung.

#### Globales Logistik- und Versandsystem

Ein System, das von einem internationalen Logistikunternehmen wie DHL oder FedEx verwendet wird.
Es besteht aus einem Netzwerk von Versandzentren, Lagern und Lieferfahrzeugen.
Die Datenverarbeitung erfolgt sowohl in lokalen Niederlassungen als auch in zentralen Rechenzentren.
Das System verwaltet Informationen zu Paketversand, Tracking, Routenoptimierung und Kundeninteraktionen.