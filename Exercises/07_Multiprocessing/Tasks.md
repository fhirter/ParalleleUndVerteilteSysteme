# Übung Prozesse

## Lernziele

- Die Studierenden wissen, wie Betriebssysteme die Ausführung von mehreren Prozessen ermöglichen.
- Die Studierenden kennen Kommunikations- und Koordinationstechniken zwischen Prozessen und Systemen

## Aufgaben

### Multiprocessing

Entwickle ein Programm, das verschiedene statistische Werte für eine Liste von Zahlen berechnet.
Das Programm soll eine Liste von zufälligen Zahlen generieren und davon den Durchschnitt, den maximalen und den
minimalen Wert berechnen.

Zum Beispiel, nehmen wir an, das Programm generiert die Zahlenreihe `90 81 78 95 79 72 85`.
Das Programm berechnet daraus:

- Der Durchschnittswert ist 82
- Der Minimalwert ist 72
- Der Maximalwert ist 95

Verwende für die Lösung [Go](https://go.dev/), dieses bietet sehr einfaches Multiprocessing.

Nutze die beliegende Vorlage aus Ausgangslage.

1. Führe die drei Berechnungen in einem Prozess aus. Miss die Zeit, die die Berechnung benötigt.
2. Führe die drei Funktionen in Goroutines aus. Miss die Zeit, die die Berechnung benötigt.
3. Erstelle eine Statistik, die die gemessenen Zeiten in Abhängigkeit der Länge der Zahlenreihe darstellt.

### Kommunikation

Optimiere die Berechnung für den Durchschnittswert aus der vorherigen Übung.

Nutze für die Kommunikation zwischen den Goroutines und dem Hauptprogramm buffered channels; Einen Channel für die
Daten, einen für die Resultate.

Buffered channels erstellst du mit `messages := make(chan string, 2)`. Der zweite Parameter von make ist die Grösse des
Buffers, d.h. die Anzahl Elemente, die im Channel gespeichert werden können.

Führe für jeden Optimierungsschritt Messungen durch, damit du die Optimierung mit Zahlen belegen kannst.

1. Recherchiere, wie Go Channels funktionieren und vergleiche mit der Beschreibung im Kapitel 3.6 (Silberschatz, 2019).
2. Schreibe das ganze Array in den Daten-Channel. Starte eine Goroutine, die die Daten aus dem Channel liest, den
   Mittelwert berechnet und das Resultat in den Resultat-Channel schreibt.
3. Unterteile die Daten in mehrere kleinere Arrays, die du in den Daten-Channel schreibst. Ändere die Goroutine so, dass
   sie läuft, bis keine Daten mehr im Daten-Channel vorhanden sind.
4. Starte nun mehrere Goroutines, die die Daten verarbeiten.
