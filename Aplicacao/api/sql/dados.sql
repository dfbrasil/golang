insert into usuarios (nome, nick, email, senha)
values
("Usuário 1", "usuario_1", "usuario_1@gmail.com","$2a$10$1Qi/5sfPiOVhHGaej0AkneGYr9v2rT/LaLLzgPk2D.sP6nTtN1fWS"),
("Usuário 2", "usuario_2", "usuario_2@gmail.com","$2a$10$1Qi/5sfPiOVhHGaej0AkneGYr9v2rT/LaLLzgPk2D.sP6nTtN1fWS"),
("Usuário 3", "usuario_3", "usuario_3@gmail.com","$2a$10$1Qi/5sfPiOVhHGaej0AkneGYr9v2rT/LaLLzgPk2D.sP6nTtN1fWS");

insert into seguidores (usuario_id, seguidor_id)
values
(1, 2),
(3, 1),
(1, 3);
