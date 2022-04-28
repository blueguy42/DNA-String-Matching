CREATE TABLE disease (
    name varchar(50) NOT NULL,
    dna varchar(100) DEFAULT NULL,
    PRIMARY KEY(name)
);

CREATE TABLE history (
    id SERIAL,
    date varchar(10) DEFAULT NULL,
    name varchar(50) DEFAULT NULL,
    disease varchar(50) DEFAULT NULL,
    result SMALLINT DEFAULT NULL,
    similarity INT NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT history_ibfk_1 FOREIGN KEY (disease) REFERENCES disease (name)
);

INSERT INTO disease (name,dna) VALUES ('cacarair','ATGGTGCACGAT'),('flu','ATGCTGACGAT');
INSERT INTO history (date,name,disease,result,similarity) VALUES ('2022-03-01','saul','flu',0,0),('2022-01-30','afan','cacarair',0,0);