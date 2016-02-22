## Übung: Key-Value Store (Cont.)

* Schreibe einen kleines Programm `kv`, mit dem Du Schlüssel-Werte Paare in einer Datei speichern und abfragen kannst
* Sichere Deine Key-Value Store Klasse so ab, dass sie parallelem Zugriff Stand hält
* Teste die Klasse mit vielen parallelen Reads und Writes
* Messe die Zeit, die Dein KV-Store braucht um einen Durchlauf von Schreiben-Speichern-Lesen durch zu führen

Installieren:
```shell
cd /path/to/go-workshow
go install session01/kv
``` 

Setzen von Werten:
```shell
kv name=Mancke vorname=Sebastian alter=42
```

Abfragen bestimmter Werte:
```shell
kv name vorname
> name=Mancke
> vorname=Sebastian
```

Abfragen aller Werte:
```shell
kv
> name=Mancke
> vorname=Sebastian
> alter=42
```