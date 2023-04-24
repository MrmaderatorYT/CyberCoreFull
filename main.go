package main

// Импортируем пакеты
/*
#include "world_generator.h"
#include <stdlib.h>
*/
import (
	"C"
	"encoding/binary"
	_ "fmt"
	"github.com/Tnze/go-mc/net"
	"log"
	"os"
)

// Packet IDs
const (
	PlayerPositionAndLookClientbound = 0x36
	JoinGame                         = 0x26
)
const (
	MaxPlayer       = 100
	ProtocolVersion = 340
)

func main() {

	// main - Функция запуска сервера
	// Запускаем сокет по адрессу 0.0.0.0:25565
	loop, err := net.ListenMC(":25565")
	// Если есть ошибка, то выводим её
	if err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
	f, err := os.Create("./world/region/r.0.0.mca")
	if err != nil {
		log.Fatalf("Ошибка при генерации мира: %v", err)
	}
	defer f.Close()

	// Записываем три чанка в файл
	for x := 0; x < 3; x++ {
		for z := 0; z < 3; z++ {
			chunkHeader := make([]byte, 8)
			binary.BigEndian.PutUint32(chunkHeader[0:4], uint32(x))
			binary.BigEndian.PutUint32(chunkHeader[4:8], uint32(z))
			//fmt.Println(len(chunkHeader), cap(chunkHeader))

			chunkData := make([]byte, 32768)
			for y := 0; y < 128; y++ {
				for z := 0; z < 16; z++ {
					for x := 0; x < 16; x++ {
						index := ((y*16)+z)*16 + x
						if y == 0 {
							chunkData[index] = 7 // Нижний слой - бедрок
						} else if y == 1 || y == 2 {
							chunkData[index] = 3 // Верхние слои - земля
						} else {
							chunkData[index] = 0 // Остальное - воздух
						}
					}
				}
			}

			_, err = f.Write(chunkHeader)
			if err != nil {
				panic(err)
			}

			_, err = f.Write(chunkData)
			if err != nil {
				panic(err)
			}
		}
	}
	// Цикл обрабатывающий входящие подключеня
	for {
		// Принимаем подключение или ждём
		connection, err := loop.Accept()
		// Если произошла ошибка - пропускаем соденение
		if err != nil {
			continue
		} // Принимаем подключение и обрабатываем его не блокируя основной поток
		go acceptConnection(connection)
	}
}
