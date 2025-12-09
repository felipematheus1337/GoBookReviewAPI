# GoBookReviewAPI

> API REST em Go para gerenciamento de livros e avaliações (reviews).  
> Projeto simples, porém estruturado — ideal para estudos, testes e como base para sistemas reais.

---

## 🚀 Visão Geral

A **GoBookReviewAPI** oferece endpoints básicos para cadastro de livros e reviews, com relacionamento 1:N entre livros e suas avaliações. Foi concebida como um projeto leve e funcional, com foco em:

- simplicidade de uso,
- validações essenciais,
- código limpo e organizado,
- boa base para aprendizado ou expansão futura.

---

## 📚 Entidades Principais

### Book
- `id` — UUID ou int
- `title` — string (obrigatório)
- `author` — string (obrigatório)
- `publishedYear` — int (opcional)
- `createdAt` — datetime
- `updatedAt` — datetime

### Review
- `id` — UUID ou int
- `bookId` — foreign key (referência a Book)
- `reviewer` — string (opcional)
- `rating` — int (1–5)
- `comment` — string (mínimo 5 caracteres)
- `createdAt` — datetime

---

## ✅ Casos de Uso / Endpoints

- Criar livro
- Listar livros
- Criar review para um livro
- Listar reviews de um livro
- *(Opcional)* Deletar review

---

## ✨ Por que utilizar este projeto?

- Arquitetura simples com relacionamento 1:N entre livros e reviews.
- Validações claras e regras de negócio mínimas — ideal para learning ou base de projeto.
- Poucos endpoints, o que facilita entendimento e manutenção.
- Excelente para praticar logging, estrutura de handlers/serviços/DTOs e padrões REST.
- Serve como base sólida para evoluções (autenticação, versionamento, documentação, etc.).

---

## 🧰 Pré-requisitos

- Go 1.x
- Banco de dados (configurado conforme seu ambiente — ex: PostgreSQL, SQLite, etc.)
- Variáveis de ambiente / arquivo `.env` se aplicável

---
