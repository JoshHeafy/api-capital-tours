CREATE DATABASE capital_tours;

CREATE TABLE
    propietarios (
        numero_documento VARCHAR(11) NOT NULL PRIMARY KEY,
        nombre_propietario VARCHAR(100) NOT NULL,
        direccion VARCHAR(150),
        referencia VARCHAR(150),
        tipo_documento int NOT NULL,
        telefono VARCHAR(20),
        email VARCHAR(50)
    );

CREATE TABLE
    vehiculos (
        numero_placa VARCHAR(7) NOT NULL PRIMARY KEY,
        marca VARCHAR(15) NOT NULL,
        modelo VARCHAR(15) NOT NULL,
        anio INT NOT NULL,
        color VARCHAR(7) NOT NULL,
        numero_serie VARCHAR(20) NOT NULL,
        numero_pasajeros INT NOT NULL DEFAULT 4,
        numero_asientos INT NOT NULL DEFAULT 5,
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
        estado INT NOT NULL default 0,
        fecha_fin VARCHAR(10) DEFAULT NULL,
        numero_flota INT NOT NULL,
        numero_placa VARCHAR(7) NOT NULL,
        CONSTRAINT fk_inscripciones_vehiculos FOREIGN KEY (numero_placa) REFERENCES vehiculos (numero_placa),
        CONSTRAINT fk_inscripciones_propietarios FOREIGN KEY (numero_documento) REFERENCES propietarios (numero_documento)
    );

CREATE TABLE
    detalle_inscripciones (
        id_detalle_inscripcion VARCHAR(36) NOT NULL PRIMARY KEY,
        fecha_pago VARCHAR(10) NOT NULL,
        years int NOT NULL,
        months int NOT NULL, --//entre 1 y 12
        importe FLOAT8 NOT NULL DEFAULT 0.0,
        numero_documento VARCHAR(12) NOT NULL,
        estado INT DEFAULT 0,
        id_inscripcion varchar(36) not null,
        CONSTRAINT fk_detalle_inscripciones_inscripciones FOREIGN KEY (id_inscripcion) REFERENCES inscripciones (id_inscripcion)
    );

CREATE TABLE
    comprobante_pago (
        id_comprobante_pago VARCHAR(36) NOT NULL PRIMARY KEY,
        numero_documento VARCHAR(11) NOT NULL,
        tipo varchar(2) NULL DEFAULT 0,
        numero_serie VARCHAR(4) NOT NULL,
        numero_comprobante VARCHAR(10) NOT NULL,
        fecha_pago VARCHAR(10) NOT NULL,
        importe FLOAT8 NOT NULL DEFAULT 0.0,
        igv FLOAT8 NOT NULL DEFAULT 0.0,
        descuento FLOAT8 NOT NULL DEFAULT 0.0,
        total FLOAT8 NOT NULL DEFAULT 0.0,
        observaciones VARCHAR(100) NOT NULL DEFAULT '',
        estado INT DEFAULT 0,
        id_inscripcion VARCHAR(36) NOT NULL,
        CONSTRAINT fk_comprobante_pago_inscripciones FOREIGN KEY (id_inscripcion) REFERENCES inscripciones (id_inscripcion)
    );

CREATE TABLE
    detalle_comprobantes (
        id_comprobante_pago VARCHAR(36) NOT NULL,
        item int not null,
        importe FLOAT8 NOT NULL DEFAULT 0.0,
        descuento FLOAT8 NOT NULL DEFAULT 0.0,
        igv FLOAT8 NOT NULL DEFAULT 0.0,
        total FLOAT8 NOT NULL DEFAULT 0.0,
        years int NOT NULL,
        months int NOT NULL,
        CONSTRAINT fk_detalle_comprobantes_comprobante_pago FOREIGN KEY (id_comprobante_pago) REFERENCES comprobante_pago (id_comprobante_pago)
    );

CREATE TABLE
    permisos (
        id_permiso VARCHAR(36) NOT NULL PRIMARY KEY,
        fecha_inicio VARCHAR(10) NOT NULL,
        fecha_fin VARCHAR(10) NOT NULL,
        id_inscripcion VARCHAR(36) NOT NULL,
        CONSTRAINT fk_permisos_inscripciones FOREIGN KEY (id_inscripcion) REFERENCES inscripciones (id_inscripcion)
    );

--web
CREATE TABLE
    solicitudes (
        id_solicitudes VARCHAR(36) NOT NULL PRIMARY KEY,
        nombre VARCHAR(100) NOT NULL,
        email VARCHAR(50) NOT NULL,
        telefono VARCHAR(15) NOT NULL,
        asunto VARCHAR(100) NOT NULL,
        mensaje VARCHAR(500) NOT NULL
    );

--movil
CREATE TABLE
    users (
        id_user VARCHAR(36) NOT NULL PRIMARY KEY,
        cargo INT DEFAULT 0,
        username VARCHAR(100) NOT NULL,
        email VARCHAR(100) NOT NULL,
        password VARCHAR(200) NOT NULL,
        numero_placa VARCHAR(36) NOT NULL,
        CONSTRAINT fk_users_vehiculos FOREIGN KEY (numero_placa) REFERENCES vehiculos (numero_placa),
        CONSTRAINT unique_email UNIQUE (email),
        CONSTRAINT unique_username UNIQUE (username)
    );

------------------------------------------------------------------------
-- Insertar propietarios
INSERT INTO
    propietarios (
        numero_documento,
        nombre_propietario,
        direccion,
        referencia,
        tipo_documento,
        telefono,
        email
    )
VALUES
    (
        '87654321',
        'alex smith',
        'Dirección 2',
        'Referencia 2',
        2,
        '987654321',
        'propietario2@example.com'
    ),
    (
        '98765432',
        'marlon wayan',
        'Dirección 3',
        'Referencia 3',
        1,
        '654321987',
        'propietario3@example.com'
    ),
    (
        '00000000',
        'supervisor',
        'supervisor',
        'ninguna',
        0,
        '000000000',
        'supervisor@capitaltours.com'
    ),
    (
        '12345678901',
        'jhon does',
        '123 Main St',
        'near parks',
        1,
        '555-1234',
        'john.doe@example.com'
    ),
    (
        '12345678',
        'matt kunshall',
        'Dirección 1',
        'Referencia 1',
        1,
        '123456789',
        'propietario1@example.com'
    ),
    (
        '78978987878',
        'mateo sape',
        'av siempreviva 123',
        'entre dos postes',
        1,
        '987564848',
        'mateo@gmail.com'
    );

-- Insertar vehículos
INSERT INTO
    vehiculos (
        numero_placa,
        marca,
        modelo,
        anio,
        color,
        numero_serie,
        numero_pasajeros,
        numero_asientos,
        observaciones,
        numero_documento
    )
VALUES
    (
        '000-000',
        'supervisor',
        'supervisor',
        2023,
        '#ffffff',
        '000000',
        4,
        5,
        'ninguna',
        '00000000'
    ),
    (
        '987-OYP',
        'Marca 1',
        'Modelo 1',
        2020,
        '#cde54e',
        'ABC123',
        4,
        5,
        'Observaciones 1',
        '12345678'
    ),
    (
        '654-RFD',
        'Marca 2',
        'Modelo 2',
        2018,
        '#ea525f',
        'DEF456',
        4,
        5,
        'Observaciones 2',
        '12345678'
    ),
    (
        '321-KJH',
        'Marca 3',
        'Modelo 3',
        2019,
        '#abece4',
        'GHI789',
        4,
        5,
        'Observaciones 3',
        '87654321'
    ),
    (
        '789-FGH',
        'Marca 4',
        'Modelo 4',
        2021,
        '#e15e6e',
        'JKL012',
        4,
        5,
        'Observaciones 4',
        '98765432'
    ),
    (
        'ABC123',
        'Toyota',
        'Camry',
        2022,
        '#FFFFFF',
        '1234567890',
        4,
        5,
        'No observations',
        '12345678901'
    ),
    (
        'ASD-589',
        'toyota',
        'toyota',
        2015,
        '#EDEBC9',
        '5821546',
        4,
        4,
        'ninguna',
        '87654321'
    ),
    (
        'FRT-586',
        'nissan',
        'nissan',
        2018,
        '#FF5254',
        '857415851',
        4,
        5,
        'ninguna',
        '98765432'
    );

-- Insertar inscripciones
INSERT INTO
    inscripciones (
        id_inscripcion,
        numero_documento,
        fecha_inicio,
        importe,
        fecha_pago,
        years,
        months,
        estado,
        fecha_fin,
        numero_flota,
        numero_placa
    )
VALUES
    (
        '24bd7d9b-c055-4d8d-b2c7-14c64299f5a6',
        '98765432',
        '13/07/2023',
        150,
        '13/08/2023',
        2023,
        1,
        1,
        NULL,
        18,
        '789-FGH'
    ),
    (
        '6f41e24a-20e1-4685-b57b-c1739a93a5a5',
        '12345678901',
        '13/07/2023',
        150,
        '13/08/2023',
        2023,
        1,
        1,
        NULL,
        37,
        'ABC123'
    ),
    (
        'cd06782e-a884-40f7-bf06-43dd57a79c85',
        '12345678',
        '01/01/2023',
        100,
        '05/01/2023',
        1,
        6,
        0,
        '31/12/2023',
        1,
        '987-OYP'
    ),
    (
        '3b9cdc6e-6f67-4345-9134-d0e39aaf279d',
        '87654321',
        '01/01/2023',
        200,
        '05/01/2023',
        1,
        6,
        0,
        '31/12/2023',
        3,
        '321-KJH'
    ),
    (
        '30e65b6e-4899-4ef9-9034-0aaac13a3f51',
        '12345678',
        '13/07/2023',
        100,
        '13/08/2023',
        2023,
        7,
        0,
        NULL,
        36,
        '654-RFD'
    ),
    (
        '500429cd-e258-43c2-b1c2-7e225e80df1b',
        '12345678',
        '01/01/2023',
        150,
        '05/01/2023',
        1,
        12,
        0,
        '31/12/2023',
        2,
        '654-RFD'
    ),
    (
        '1',
        '12345678901',
        '01/01/2023',
        100,
        '01/01/2023',
        1,
        6,
        1,
        '30/06/2023',
        10,
        'ABC123'
    ),
    (
        '2',
        '12345678901',
        '01/01/2023',
        100,
        '01/02/2023',
        2023,
        6,
        1,
        NULL,
        25,
        'ABC123'
    ),
    (
        '13dbb498-9414-415a-942b-e0804c06b4ce',
        '98765432',
        '01/01/2023',
        250,
        '05/01/2023',
        1,
        12,
        1,
        '31/12/2023',
        4,
        '789-FGH'
    ),
    (
        'f54e779b-a8ee-4527-bc55-b93b57f361a3',
        '12345678901',
        '01/01/2023',
        100,
        '01/02/2023',
        2023,
        6,
        1,
        NULL,
        25,
        'ABC123'
    ),
    (
        '22c00ad3-1652-4ac9-bc6f-3ec59bc0a6d5',
        '12345678',
        '13/07/2023',
        100,
        '13/08/2023',
        2023,
        7,
        1,
        NULL,
        36,
        '654-RFD'
    );

-- Insertar detalle_inscripciones
INSERT INTO
    detalle_inscripciones (
        id_detalle_inscripcion,
        fecha_pago,
        years,
        months,
        importe,
        numero_documento,
        estado,
        id_inscripcion
    )
VALUES
    (
        '1',
        '01/01/2023',
        1,
        6,
        100,
        '12345678901',
        0,
        '1'
    ),
    (
        '9f4baa31-826c-4330-81a9-81966a89769d',
        '05/01/2023',
        1,
        6,
        100,
        '12345678',
        1,
        'cd06782e-a884-40f7-bf06-43dd57a79c85'
    ),
    (
        '04c4fb71-35a7-45f8-9bfc-d2c4aab22d0e',
        '05/01/2023',
        1,
        12,
        150,
        '12345678',
        1,
        '500429cd-e258-43c2-b1c2-7e225e80df1b'
    ),
    (
        '2e800179-55b9-4819-81da-4ce46b6f429c',
        '05/01/2023',
        1,
        6,
        200,
        '87654321',
        1,
        '3b9cdc6e-6f67-4345-9134-d0e39aaf279d'
    ),
    (
        '2433d8b7-f4aa-47f7-9fd9-b7687dfe9c5a',
        '05/01/2023',
        1,
        12,
        250,
        '98765432',
        1,
        '13dbb498-9414-415a-942b-e0804c06b4ce'
    );

-- Insertar comprobante_pago
INSERT INTO
    comprobante_pago (
        id_comprobante_pago,
        numero_documento,
        tipo,
        numero_serie,
        numero_comprobante,
        fecha_pago,
        importe,
        igv,
        descuento,
        total,
        observaciones,
        estado,
        id_inscripcion
    )
values
    (
        '373cb35e-fcb7-42c0-baac-52c9e0db6f20',
        '12345678',
        '01',
        '1234',
        '00001',
        '05/01/2023',
        100,
        18,
        0,
        118,
        'Observaciones 1',
        1,
        'cd06782e-a884-40f7-bf06-43dd57a79c85'
    ),
    (
        '48ea0c15-50a4-40be-8bc2-374fd974ccdc',
        '12345678',
        '01',
        '1234',
        '00002',
        '05/01/2023',
        150,
        27,
        0,
        177,
        'Observaciones 2',
        1,
        '500429cd-e258-43c2-b1c2-7e225e80df1b'
    ),
    (
        'c9726412-6051-4920-996d-0451758132ea',
        '87654321',
        '01',
        '1234',
        '00003',
        '05/01/2023',
        200,
        36,
        0,
        236,
        'Observaciones 3',
        1,
        '3b9cdc6e-6f67-4345-9134-d0e39aaf279d'
    ),
    (
        '63d64890-e643-43d0-97fd-a68d159fd596',
        '98765432',
        '01',
        '1234',
        '00004',
        '05/01/2023',
        250,
        45,
        0,
        295,
        'Observaciones 4',
        1,
        '13dbb498-9414-415a-942b-e0804c06b4ce'
    ),
    (
        'a44599e8-302d-4e11-bf2a-8cfe9e8a3b92',
        '98765432',
        '01',
        '0001',
        '00001',
        '13/08/2023',
        150,
        1,
        0,
        151,
        '',
        1,
        '24bd7d9b-c055-4d8d-b2c7-14c64299f5a6'
    ),
    (
        '20adc410-35d7-41f8-b4cc-44c3c6a553b2',
        '12345678901',
        '01',
        '005',
        '00002',
        '13/08/2023',
        150,
        2,
        0,
        150,
        '',
        1,
        '6f41e24a-20e1-4685-b57b-c1739a93a5a5'
    );

-- Insertar detalle_comprobantes
INSERT INTO
    detalle_comprobantes (
        id_comprobante_pago,
        item,
        importe,
        descuento,
        igv,
        total,
        years,
        months
    )
VALUES
    (
        '373cb35e-fcb7-42c0-baac-52c9e0db6f20',
        1,
        100,
        0,
        18,
        118,
        1,
        6
    ),
    (
        '48ea0c15-50a4-40be-8bc2-374fd974ccdc',
        1,
        150,
        0,
        27,
        177,
        1,
        12
    ),
    (
        'c9726412-6051-4920-996d-0451758132ea',
        1,
        200,
        0,
        36,
        236,
        1,
        6
    ),
    (
        '63d64890-e643-43d0-97fd-a68d159fd596',
        1,
        250,
        0,
        45,
        295,
        1,
        12
    );

-- Insertar permisos
INSERT INTO
    permisos (
        id_permiso,
        fecha_inicio,
        fecha_fin,
        id_inscripcion
    )
VALUES
    (
        '8aaf0ca4-f473-481f-8418-560bc1949852',
        '01/01/2023',
        '31/12/2023',
        'cd06782e-a884-40f7-bf06-43dd57a79c85'
    ),
    (
        'dc2844dd-9014-4d85-9326-4044872b09aa',
        '01/01/2023',
        '31/12/2023',
        '500429cd-e258-43c2-b1c2-7e225e80df1b'
    ),
    (
        '4dbb7eb7-c3bf-491f-b963-a243d0b6093f',
        '01/01/2023',
        '31/12/2023',
        '3b9cdc6e-6f67-4345-9134-d0e39aaf279d'
    ),
    (
        'e4c12eee-afe8-49c4-80e6-fb5d85cc3e58',
        '01/01/2023',
        '31/12/2023',
        '13dbb498-9414-415a-942b-e0804c06b4ce'
    );

-- Insertar solicitudes
INSERT INTO
    solicitudes (
        id_solicitudes,
        nombre,
        email,
        telefono,
        asunto,
        mensaje
    )
VALUES
    (
        '1c371489-d193-4b72-aa43-639871d5a3ce',
        'Usuario 1',
        'usuario1@example.com',
        '123456789',
        'Asunto 1',
        'Mensaje 1'
    ),
    (
        '5a5af779-2845-4204-a637-05b29b9544a3',
        'Usuario 2',
        'usuario2@example.com',
        '987654321',
        'Asunto 2',
        'Mensaje 2'
    ),
    (
        '9b34d6a9-402b-4937-98bb-599a85a423c9',
        'Usuario 3',
        'usuario3@example.com',
        '654321987',
        'Asunto 3',
        'Mensaje 3'
    );

-- Insertar users
INSERT INTO
    users (
        id_user,
        cargo,
        username,
        email,
        password,
        numero_placa
    )
VALUES
    (
        'f845678a-6c9c-4e77-94a1-03ca19dbdf49',
        0,
        'supervisor',
        'supervisor@capitaltours.com',
        '$2a$10$X.QKWoH08fwR1k2d0nDSh.hefZ5AbNKSLr4xi2TCp5XO7rBNvojtm',
        '000-000'
    );