#pragma once
#include <string>
#include <WS2tcpip.h>
#include <msclr\marshal_cppstd.h>
#include <iostream>

namespace ClientGUI {

	using namespace System;
	using namespace System::ComponentModel;
	using namespace System::Collections;
	using namespace System::Windows::Forms;
	using namespace System::Data;
	using namespace System::Drawing;

	using namespace std;

	/// <summary>
	/// Riepilogo per MyForm
	/// </summary>
	public ref class MyForm : public System::Windows::Forms::Form
	{
	private:
		Timer^ MyTimer;

	public:
		String^ MyInput;

	public:
		SOCKET MySock;

	public:
		MyForm(void)
		{
			InitializeComponent();
			//
			//TODO: aggiungere qui il codice del costruttore.
			//

			MySock = 0;

			MyTimer = gcnew System::Windows::Forms::Timer();
			MyTimer->Interval = 3000;
			MyTimer->Tick += gcnew System::EventHandler(this, &MyForm::CallMe_OnTick);
			MyTimer->Enabled = true;
		}

	protected:
		/// <summary>
		/// Pulire le risorse in uso.
		/// </summary>
		~MyForm()
		{
			if (components)
			{
				delete components;
			}
		}
	private: System::Windows::Forms::Label^ label1;
	private: System::Windows::Forms::TextBox^ textBox1;
	private: System::Windows::Forms::Button^ button1;
	protected:

	private:
		/// <summary>
		/// Variabile di progettazione necessaria.
		/// </summary>
		System::ComponentModel::Container^ components;

#pragma region Windows Form Designer generated code
		/// <summary>
		/// Metodo necessario per il supporto della finestra di progettazione. Non modificare
		/// il contenuto del metodo con l'editor di codice.
		/// </summary>
		void InitializeComponent(void)
		{
			this->label1 = (gcnew System::Windows::Forms::Label());
			this->textBox1 = (gcnew System::Windows::Forms::TextBox());
			this->button1 = (gcnew System::Windows::Forms::Button());
			this->SuspendLayout();
			// 
			// label1
			// 
			this->label1->AutoSize = true;
			this->label1->BackColor = System::Drawing::SystemColors::ButtonFace;
			this->label1->Font = (gcnew System::Drawing::Font(L"Microsoft Sans Serif", 8.25F, System::Drawing::FontStyle::Bold, System::Drawing::GraphicsUnit::Point,
				static_cast<System::Byte>(0)));
			this->label1->ForeColor = System::Drawing::SystemColors::ControlText;
			this->label1->Location = System::Drawing::Point(57, 428);
			this->label1->Name = L"label1";
			this->label1->Size = System::Drawing::Size(127, 13);
			this->label1->TabIndex = 0;
			this->label1->Text = L"Messaggio da inviare";
			this->label1->Click += gcnew System::EventHandler(this, &MyForm::label1_Click);
			// 
			// textBox1
			// 
			this->textBox1->Location = System::Drawing::Point(190, 428);
			this->textBox1->Multiline = true;
			this->textBox1->Name = L"textBox1";
			this->textBox1->Size = System::Drawing::Size(205, 45);
			this->textBox1->TabIndex = 1;
			// 
			// button1
			// 
			this->button1->Location = System::Drawing::Point(190, 479);
			this->button1->Name = L"button1";
			this->button1->Size = System::Drawing::Size(75, 26);
			this->button1->TabIndex = 2;
			this->button1->Text = L"Invia";
			this->button1->UseVisualStyleBackColor = true;
			this->button1->Click += gcnew System::EventHandler(this, &MyForm::button1_Click);
			// 
			// MyForm
			// 
			this->AutoScaleDimensions = System::Drawing::SizeF(6, 13);
			this->AutoScaleMode = System::Windows::Forms::AutoScaleMode::Font;
			this->ClientSize = System::Drawing::Size(670, 531);
			this->Controls->Add(this->button1);
			this->Controls->Add(this->textBox1);
			this->Controls->Add(this->label1);
			this->Name = L"MyForm";
			this->Text = L"MyForm";
			this->Load += gcnew System::EventHandler(this, &MyForm::MyForm_Load);
			this->ResumeLayout(false);
			this->PerformLayout();

		}
#pragma endregion

	private: 
		System::Void MyForm_Load(System::Object^ sender, System::EventArgs^ e)
		{
		}

		System::Void label1_Click(System::Object^ sender, System::EventArgs^ e)
		{
		}

		System::Void button1_Click(System::Object^ sender, System::EventArgs^ e) 
		{
			MyInput = textBox1->Text;
			//MessageBox::Show(MyInput);
		}

		void CallMe_OnTick(System::Object^ sender, System::EventArgs^ e)
		{
			MyTimer->Interval = 1000;

			if (MySock != 0 && String::IsNullOrEmpty(MyInput) == false)
			{
				char buf[4096];
				string localInput;

				//do
				//{
					msclr::interop::marshal_context context;
					localInput = context.marshal_as<string>(MyInput);

					textBox1->Text += "";
					MyInput = "";

					if (localInput.size() > 0)		// Make sure the user has typed in something
					{
						// Send the text
						int sendResult = send(MySock, localInput.c_str(), localInput.size() + 1, 0);
						if (sendResult != SOCKET_ERROR)
						{
							// Wait for response
							ZeroMemory(buf, 4096);
							int bytesReceived = recv(MySock, buf, 4096, 0);
							if (bytesReceived > 0)
							{
								// Echo response to console
								cout << "SERVER> " << string(buf, 0, bytesReceived) << endl;
								textBox1->Text += bytesReceived;
							}
						}
					}

				//} while (localInput.size() > 0);
			}
		}
	private: System::Void label2_Click(System::Object^ sender, System::EventArgs^ e) {
	}
};
}