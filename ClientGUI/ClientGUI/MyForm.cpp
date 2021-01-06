#include "MyForm.h"
#include <iostream>
#include <string>
#include <WS2tcpip.h>

#pragma comment(lib, "ws2_32.lib")

using namespace std;
using namespace ClientGUI;
using namespace System;
using namespace System::Windows::Forms;
[STAThreadAttribute]

void main() {

	string ipAddress = "127.0.0.1";			// IP Address of the server
	int port = 8080;						// Listening port # on the server

	// Initialize WinSock
	WSAData data;
	WORD ver = MAKEWORD(2, 2);
	int wsResult = WSAStartup(ver, &data);
	if (wsResult != 0)
	{
		cerr << "Can't start Winsock, Err #" << wsResult << endl;
		return;
	}

	// Create socket
	SOCKET sock = socket(AF_INET, SOCK_STREAM, 0);
	if (sock == INVALID_SOCKET)
	{
		cerr << "Can't create socket, Err #" << WSAGetLastError() << endl;
		WSACleanup();
		return;
	}

	// Fill in a hint structure
	sockaddr_in hint;
	hint.sin_family = AF_INET;
	hint.sin_port = htons(port);
	inet_pton(AF_INET, ipAddress.c_str(), &hint.sin_addr);

	// Connect to server
	int connResult = connect(sock, (sockaddr*)&hint, sizeof(hint));
	if (connResult == SOCKET_ERROR)
	{
		cerr << "Can't connect to server, Err #" << WSAGetLastError() << endl;
		closesocket(sock);
		WSACleanup();
		return;
	}

	Application::SetCompatibleTextRenderingDefault(false);
	Application::EnableVisualStyles();
	ClientGUI::MyForm frm;
	frm.MySock = sock;
	Application::Run(% frm);
	
	/*
			// Do-while loop to send and receive data
			char buf[4096];
			string userInput;
			//string userInput = frm.userInput;
			//userInput = testo;
			do
			{
				// Prompt the user for some text
				cout << "> ";
				getline(cin, userInput);

				if (userInput.size() > 0)		// Make sure the user has typed in something
				{
					// Send the text
					int sendResult = send(sock, userInput.c_str(), userInput.size() + 1, 0);
					if (sendResult != SOCKET_ERROR)
					{
						// Wait for response
						ZeroMemory(buf, 4096);
						int bytesReceived = recv(sock, buf, 4096, 0);
						if (bytesReceived > 0)
						{
							// Echo response to console
							cout << "SERVER> " << string(buf, 0, bytesReceived) << endl;
						}
					}
				}

			} while (userInput.size() > 0);
	*/


	// Gracefully close down everything
	closesocket(sock);
	WSACleanup();
}

// Per eseguire il programma: CTRL+F5 oppure Debug > Avvia senza eseguire debug
// Per eseguire il debug del programma: F5 oppure Debug > Avvia debug

// Suggerimenti per iniziare: 
//   1. Usare la finestra Esplora soluzioni per aggiungere/gestire i file
//   2. Usare la finestra Team Explorer per connettersi al controllo del codice sorgente
//   3. Usare la finestra di output per visualizzare l'output di compilazione e altri messaggi
//   4. Usare la finestra Elenco errori per visualizzare gli errori
//   5. Passare a Progetto > Aggiungi nuovo elemento per creare nuovi file di codice oppure a Progetto > Aggiungi elemento esistente per aggiungere file di codice esistenti al progetto
//   6. Per aprire di nuovo questo progetto in futuro, passare a File > Apri > Progetto e selezionare il file con estensione sln
