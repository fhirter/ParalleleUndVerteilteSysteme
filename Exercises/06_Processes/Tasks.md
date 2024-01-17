# Übung Prozesse

Halte dein Vorgehen, Überlegungen und Erkenntnisse in einem Dokument fest.
Schicke dieses am Ende der Unterrichtseinheit dem Dozenten.

Organisiert euch selbstständig in der Klasse, wie ihr die Aufgabe Lösen wollt.
Stellt sicher, dass alle verstehen, was gemacht wurde.

## Lernziele

- Die Studierenden wissen, wie Betriebssysteme die Ausführung von mehreren Prozessen ermöglichen.
- Die Studierenden kennen Kommunikations- und Koordinationstechniken zwischen Prozessen und Systemen

## Aufgabenstellung

Halte die Erkenntnisse in einem [Markdown-Dokument](https://www.markdownguide.org/) fest.

Stelle sicher, dass deine Untersuchungen vollständig nachvollziehbar sind!

**Schicke das Dokument nach dem Unterricht an den Dozenten.**

### Systemanalyse

Verbinde dich mit den separat kommunizierten Zugangsdaten mit dem Linux-System.

Untersuche anhand der laufenden Prozesse, wozu das System dient.
Identifiziere die wichtigsten Prozesse.

Untersuche diese identifizierten Prozesse mit `ps` und `htop` indem du folgende Fragen beantwortest.
Beschränke dich dabei nicht nur auf die reine beantwortung der Frage, sondern erkläre die Frage an sich sowie deinen 
Lösungsprozess.

- Welche Prozesse nutzen wie viel Speicher?
- Welche Prozesse nutzen wie viel CPU?
- Welche Prozesse haben wie viel CPU Zeit verwendet?
- Welchen Ausführungsstatus haben die Prozesse?
- Welche Scheduling Priorität haben die Prozesse?
- Wie viele Prozesse starten die einzelnen Applikationen?
- Welche Prozesse nutzen mehrere Threads?

## Concurrency

Ein System "A" veröffentlicht bei jedem Datenbankzugriff die geänderten Daten in ein Publisher Subscriber System (
Message Broker).
Es werden bis zu 100 Nachrichten pro Sekunde veröffentlicht.
Die Nachrichten beziehen sich jeweils auf eine Entität, welche mit einer GUID identifiziert wird.

Das System "B" erhält von diesem Message Broker die Nachrichten und schreib diese in eine Datenbank.
In der Datenbank ist die GUID als primary Key definiert.
Beim Schreiben muss die Datenbank zuerst abgefragt werden, ob der Eintrag schon existiert und anschliessend der Eintrag
aktualisiert oder erstellt werden.

![](SentMessages.png)

Der Message Broker wartet mit der Zustellung nicht, bis die vorherige Nachricht bestätigt wurde.
Das Empfangssystem verarbeitet die Nachrichten in mehreren Prozessen.

## Problemstellung

Es tauchen selten aber regelmässig «Duplicate Entry» Fehler auf.

- Was könnte die Ursache für diese Fehler sein? Skizziere mit einem Aktivitätsdiagramm.
  Verwende dazu PlantUML oder Mermaid.
- Wie könnte das Problem vermieden werden?
