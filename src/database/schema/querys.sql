-- CREATE DATABASE capital_tours;
CREATE TABLE
    propietarios (
        numero_documento VARCHAR(11) NOT NULL PRIMARY KEY,
        nombre_propietario VARCHAR(200) NOT NULL,
        direccion VARCHAR(150) NOT NULL,
        referencia VARCHAR(150),
        tipo_documento int NOT NULL,
        telefono VARCHAR(9),
        email VARCHAR(150)
    );

CREATE TABLE
    vehiculos (
        numero_placa VARCHAR(7) NOT NULL PRIMARY KEY,
        marca VARCHAR(15) NOT NULL,
        modelo VARCHAR(50) NOT NULL,
        anio INT NOT NULL,
        color VARCHAR(7) NOT NULL,
        numero_serie VARCHAR(17) NOT NULL,
        numero_pasajeros INT NOT NULL, --default 4
        numero_asientos INT NOT NULL, --default 5
        observaciones VARCHAR(100),
        numero_documento VARCHAR(11) NOT NULL,
        CONSTRAINT fk_vehiculos_propietarios FOREIGN KEY (numero_documento) REFERENCES propietarios (numero_documento)
    );

CREATE TABLE
    inscripciones (
        id_inscripcion VARCHAR(36) NOT NULL PRIMARY KEY,
        numero_documento VARCHAR(11) NOT NULL,
        fecha_inicio VARCHAR(10) NOT NULL,
        importe FLOAT8 NOT NULL DEFAULT 0.0,
        fecha_pago VARCHAR(10) NOT NULL,
        years int NOT NULL,
        months int NOT NULL,
        estado INT NOT NULL default 0, -- 0: inactivo, 1: activo
        fecha_fin VARCHAR(10) DEFAULT NULL, -- si deja de ser null entonces la inscripcion habra terminado
        numero_flota INT NOT NULL,
        numero_placa VARCHAR(7) NOT NULL,
        CONSTRAINT fk_inscripciones_vehiculos FOREIGN KEY (numero_placa) REFERENCES vehiculos (numero_placa),
        CONSTRAINT unique_flota UNIQUE (numero_flota)
    );

CREATE TABLE
    detalle_inscripciones (
        id_detalle_inscripcion VARCHAR(36) NOT NULL PRIMARY KEY,
        fecha_pago VARCHAR(10) NOT NULL,
        years int NOT NULL,
        months int NOT NULL,
        importe FLOAT8 NOT NULL DEFAULT 0.0,
        numero_documento VARCHAR(11) NOT NULL,
        estado INT DEFAULT 0, -- 0: inactivo, 1: activo
        id_inscripcion varchar(36) not null,
        CONSTRAINT fk_detalle_inscripciones_inscripciones FOREIGN KEY (id_inscripcion) REFERENCES inscripciones (id_inscripcion)
    );

CREATE TABLE
    comprobante_pago (
        id_comprobante_pago VARCHAR(36) NOT NULL PRIMARY KEY,
        numero_documento VARCHAR(11) NOT NULL,
        tipo varchar(2) NOT NULL,
        numero_serie VARCHAR(4) NOT NULL,
        numero_comprobante VARCHAR(10) NOT NULL,
        fecha_pago VARCHAR(10) NOT NULL,
        importe FLOAT8 NOT NULL DEFAULT 0.0,
        igv FLOAT8 NOT NULL DEFAULT 0.0,
        descuento FLOAT8 NOT NULL DEFAULT 0.0,
        total FLOAT8 NOT NULL DEFAULT 0.0,
        observaciones VARCHAR(150) NOT NULL DEFAULT '',
        estado INT DEFAULT 1, -- 0: inactivo, 1: activo
        id_inscripcion VARCHAR(36) NOT NULL,
        CONSTRAINT fk_comprobante_pago_inscripciones FOREIGN KEY (id_inscripcion) REFERENCES inscripciones (id_inscripcion)
    );

CREATE TABLE
    detalle_comprobantes (
        id_comprobante_pago VARCHAR(36) NOT NULL, -- directamente amarrado a comprobante
        item serial not null, -- para borrar el registro
        importe FLOAT8 NOT NULL DEFAULT 0.0,
        igv FLOAT8 NOT NULL DEFAULT 0.0,
        descuento FLOAT8 NOT NULL DEFAULT 0.0,
        total FLOAT8 NOT NULL DEFAULT 0.0,
        years int NOT NULL,
        months int NOT NULL,
        CONSTRAINT fk_detalle_comprobantes_comprobante_pago FOREIGN KEY (id_comprobante_pago) REFERENCES comprobante_pago (id_comprobante_pago)
    );

-- CREATE TABLE
--     permisos (
--         id_permiso VARCHAR(36) NOT NULL PRIMARY KEY,
--         fecha_inicio VARCHAR(10) NOT NULL,
--         fecha_fin VARCHAR(10) NOT NULL,
--         id_inscripcion VARCHAR(36) NOT NULL,
--         CONSTRAINT fk_permisos_inscripciones FOREIGN KEY (id_inscripcion) REFERENCES inscripciones (id_inscripcion)
--     );
--web
CREATE TABLE
    solicitudes (
        id_solicitud VARCHAR(36) NOT NULL PRIMARY KEY,
        nombre VARCHAR(200) NOT NULL,
        email VARCHAR(100) NOT NULL,
        telefono VARCHAR(9) NOT NULL,
        asunto VARCHAR(150) NOT NULL,
        mensaje VARCHAR(500) NOT NULL,
        fecha_solicitud TIMESTAMP DEFAULT NOW (),
        leido INT NOT NULL DEFAULT 0
    );

--movil
CREATE TABLE
    users_mobile (
        id_user VARCHAR(36) NOT NULL PRIMARY KEY,
        cargo INT DEFAULT 0,
        email VARCHAR(150) NOT NULL,
        password VARCHAR(200) NOT NULL,
        numero_placa VARCHAR(7) NOT NULL,
        CONSTRAINT fk_users_vehiculos FOREIGN KEY (numero_placa) REFERENCES vehiculos (numero_placa),
        CONSTRAINT unique_email UNIQUE (email),
        CONSTRAINT unique_placa UNIQUE (numero_placa)
    );

CREATE TABLE
    users_admin (
        id_user_admin VARCHAR(36) NOT NULL PRIMARY KEY,
        cargo INT DEFAULT 0,
        username VARCHAR(100) NOT NULL,
        nombre VARCHAR(100) NOT NULL,
        apellidos VARCHAR(100) NOT NULL,
        id_img VARCHAR,
        email VARCHAR(150) NOT NULL,
        password VARCHAR(200) NOT NULL,
        CONSTRAINT unique_username_admin UNIQUE (username)
    );

CREATE TABLE
    locations (
        id_location VARCHAR(36) NOT NULL PRIMARY KEY,
        nombre VARCHAR(150) NOT NULL,
        latitud double precision,
        longitud double precision,
        numero_placa VARCHAR(7) NOT NULL,
        numero_flota INT NOT NULL,
    );

------------------------------------------------------------------------
-- -- Insertar propietarios
-- INSERT INTO
--     propietarios (
--         numero_documento,
--         nombre_propietario,
--         direccion,
--         referencia,
--         tipo_documento,
--         telefono,
--         email
--     )
-- VALUES
--     (
--         '87654321',
--         'alex smith',
--         'Dirección 1',
--         'Referencia 1',
--         1,
--         '987654321',
--         'propietario2@example.com'
--     ),
--     (
--         '98765432',
--         'marlon wayan',
--         'Dirección 2',
--         'Referencia 2',
--         1,
--         '654321987',
--         'propietario3@example.com'
--     ),
--     (
--         '12345678',
--         'matt kunshall',
--         'Dirección 3',
--         'Referencia 3',
--         1,
--         '123456789',
--         'propietario1@example.com'
--     ),
--     (
--         '12345678901',
--         'jhon does',
--         '123 Main St',
--         'near parks',
--         1,
--         '555-1234',
--         'john.doe@example.com'
--     ),
--     (
--         '78978987878',
--         'mateo sape',
--         'av siempreviva 123',
--         'entre dos postes',
--         1,
--         '987564848',
--         'mateo@gmail.com'
--     );
-- -- Insertar vehículos
-- INSERT INTO
--     vehiculos (
--         numero_placa,
--         marca,
--         modelo,
--         anio,
--         color,
--         numero_serie,
--         numero_pasajeros,
--         numero_asientos,
--         observaciones,
--         numero_documento
--     )
-- VALUES
--     (
--         '987-OYP',
--         'Marca 1',
--         'Modelo 1',
--         2020,
--         '#cde54e',
--         'ABC123',
--         4,
--         5,
--         'Observaciones 1',
--         '12345678'
--     ),
--     (
--         '654-RFD',
--         'Marca 2',
--         'Modelo 2',
--         2018,
--         '#ea525f',
--         'DEF456',
--         4,
--         5,
--         'Observaciones 2',
--         '12345678'
--     ),
--     (
--         '321-KJH',
--         'Marca 3',
--         'Modelo 3',
--         2019,
--         '#abece4',
--         'GHI789',
--         4,
--         5,
--         'Observaciones 3',
--         '87654321'
--     ),
--     (
--         '789-FGH',
--         'Marca 4',
--         'Modelo 4',
--         2021,
--         '#e15e6e',
--         'JKL012',
--         4,
--         5,
--         'Observaciones 4',
--         '98765432'
--     ),
--     (
--         'ABC123',
--         'Toyota',
--         'Camry',
--         2022,
--         '#FFFFFF',
--         '1234567890',
--         4,
--         5,
--         'No observations',
--         '12345678901'
--     ),
--     (
--         'ASD-589',
--         'toyota',
--         'toyota',
--         2015,
--         '#EDEBC9',
--         '5821546',
--         4,
--         4,
--         'ninguna',
--         '87654321'
--     ),
--     (
--         'FRT-586',
--         'nissan',
--         'nissan',
--         2018,
--         '#FF5254',
--         '857415851',
--         4,
--         5,
--         'ninguna',
--         '98765432'
--     );
-- -- Insertar inscripciones
-- INSERT INTO
--     inscripciones (
--         id_inscripcion,
--         numero_documento,
--         fecha_inicio,
--         importe,
--         fecha_pago,
--         years,
--         months,
--         estado,
--         fecha_fin,
--         numero_flota,
--         numero_placa
--     )
-- VALUES
--     (
--         'c0a7b6b4-cb52-4338-b2c9-344ee783ee3a',
--         '87654321',
--         '01/08/2023',
--         100,
--         '01/09/2023',
--         2023,
--         8,
--         1,
--         NULL,
--         36,
--         '321-KJH'
--     );
-- -- Insertar detalle_inscripciones
-- INSERT INTO
--     detalle_inscripciones (
--         id_detalle_inscripcion,
--         fecha_pago,
--         years,
--         months,
--         importe,
--         numero_documento,
--         estado,
--         id_inscripcion
--     )
-- VALUES
--     (
--         '616ac05b-be52-4ab8-8d02-fa28e4d4b70a',
--         '01/09/2023',
--         2023,
--         8,
--         100,
--         '87654321',
--         1,
--         'c0a7b6b4-cb52-4338-b2c9-344ee783ee3a'
--     );
-- -- Insertar comprobantes
-- INSERT INTO
--     comprobante_pago (
--         id_comprobante_pago,
--         numero_documento,
--         tipo,
--         numero_serie,
--         numero_comprobante,
--         fecha_pago,
--         importe,
--         igv,
--         descuento,
--         total,
--         observaciones,
--         estado,
--         id_inscripcion
--     )
-- VALUES
--     (
--         '415bc8fb-b811-45d4-944d-2b84be0c49c3',
--         '87654321',
--         '01',
--         '0001',
--         '0000000001',
--         '02/08/2023',
--         150,
--         0.18,
--         0,
--         177,
--         'ninguna',
--         1,
--         'c0a7b6b4-cb52-4338-b2c9-344ee783ee3a'
--     );
-- -- Insertar detalle_comprobantes
-- INSERT INTO
--     detalle_comprobantes (
--         id_comprobante_pago,
--         item,
--         importe,
--         igv,
--         descuento,
--         total,
--         years,
--         months
--     )
-- VALUES
--     (
--         '415bc8fb-b811-45d4-944d-2b84be0c49c3',
--         3,
--         150,
--         0.18,
--         0,
--         177,
--         2023,
--         8
--     );
-- -- Insertar user_admin
-- INSERT INTO
--     users_admin (
--         id_user_admin,
--         cargo,
--         username,
--         nombre,
--         apellidos,
--         id_img,
--         email,
--         password
--     )
-- VALUES
--     (
--         'bd4b2fd5-e250-4872-9faa-a7c48de0d65f',
--         0,
--         'supervisor',
--         'josh',
--         'cordova canchanya',
--         NULL,
--         'joshar456@gmail.com',
--         '$2a$10$P9CxqO3EgE0ftQL2Hpla7endolsLLVMjuG1MN6sllvwo2Ko2knIbG'
--     );
-- -- Insertar solicitudes
-- INSERT INTO
--     solicitudes (
--         id_solicitudes,
--         nombre,
--         email,
--         telefono,
--         asunto,
--         mensaje
--     )
-- VALUES
--     (
--         '1c371489-d193-4b72-aa43-639871d5a3ce',
--         'Usuario 1',
--         'usuario1@example.com',
--         '123456789',
--         'Asunto 1',
--         'Mensaje 1'
--     ),
--     (
--         '5a5af779-2845-4204-a637-05b29b9544a3',
--         'Usuario 2',
--         'usuario2@example.com',
--         '987654321',
--         'Asunto 2',
--         'Mensaje 2'
--     ),
--     (
--         '9b34d6a9-402b-4937-98bb-599a85a423c9',
--         'Usuario 3',
--         'usuario3@example.com',
--         '654321987',
--         'Asunto 3',
--         'Mensaje 3'
--     );