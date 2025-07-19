# GitHub Discord Bot API

## 📝 Descrição
API desenvolvida em Go para integração entre GitHub e Discord, permitindo notificações automáticas de eventos de Pull Requests no Discord.

## 🚀 Tecnologias

- Go 1.24.5
- PostgreSQL 15.13
- Gin Web Framework
- GORM (ORM)

## 📋 Pré-requisitos

- Go 1.24.5 ou superior
- PostgreSQL 15.13
- Configurações de ambiente (ver seção de configuração)



## 🏃 Executando o Projeto

1. Inicie o servidor:
```bash
 go run main.go
```
2. Configure o webhook no GitHub:
    - Vá para as configurações do seu repositório
    - Adicione um novo webhook
    - Configure a URL: `http://seu-servidor/webhook/github`
    - Selecione "application/json" como Content-Type
    - Escolha os eventos que deseja monitorar (Pull Requests)

## 🔐 Segurança

- Implemente autenticação para os endpoints conforme necessário
- Utilize HTTPS em produção
- Valide os payloads do webhook do GitHub

## 🤝 Contribuindo

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 📄 Licença

Este projeto está sob a licença MIT - veja o arquivo [LICENSE.md](LICENSE.md) para mais detalhes.

## ✒️ Autores

* **Julio Filizzola** - *Desenvolvimento inicial* - [juliofilizzola](https://github.com/juliofilizzola)

## 📞 Suporte

Para suporte, abra uma issue no repositório ou entre em contato através do [GitHub](https://github.com/juliofilizzola).