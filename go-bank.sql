COMMENT ON DATABASE "bankAccountApp" IS 'Le schéma qui contient les données de l''application de gestion de compte de moi';

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;
COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';
CREATE TABLE bank (
    bankid integer NOT NULL,
    name character varying(50) NOT NULL
);
CREATE SEQUENCE bank_bankid_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE bank_bankid_seq OWNER TO "bankAccountApp";

ALTER SEQUENCE bank_bankid_seq OWNED BY bank.bankid;

CREATE TABLE bankaccount (
    bankaccountid integer NOT NULL,
    accountnumber character varying(50) NOT NULL,
    userid bigint NOT NULL,
    bankid bigint NOT NULL
);


ALTER TABLE bankaccount OWNER TO "bankAccountApp";

CREATE SEQUENCE bankaccount_bankaccountid_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE bankaccount_bankaccountid_seq OWNER TO "bankAccountApp";
ALTER SEQUENCE bankaccount_bankaccountid_seq OWNED BY bankaccount.bankaccountid;
CREATE TABLE items (
    id integer,
    sub_id integer,
    name character varying(20)
);

CREATE TABLE transaction (
    transactionid integer NOT NULL,
    transaction_type_id bigint NOT NULL,
    description character varying(250) NOT NULL,
    posteddate date NOT NULL,
    userdate date NOT NULL,
    fiid character varying(20) NOT NULL,
    amount double precision,
    bankaccountid bigint
);

ALTER TABLE transaction OWNER TO "bankAccountApp";
CREATE SEQUENCE transaction_transactionid_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE transaction_type (
    transaction_type_id integer NOT NULL,
    transaction_name character varying(20) NOT NULL
);

COMMENT ON TABLE transaction_type IS 'this table contains all type of transaction debit / credit';
CREATE SEQUENCE transaction_type_transaction_type_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE users (
    userid integer NOT NULL,
    nom character varying(50) NOT NULL,
    prenom character varying(50) NOT NULL,
    pseudo character varying(50) NOT NULL,
    email character varying(50) NOT NULL,
    pwd character varying(256) NOT NULL,
    salt character varying(64) NOT NULL,
);

CREATE SEQUENCE users_userid_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER TABLE ONLY bank ALTER COLUMN bankid SET DEFAULT nextval('bank_bankid_seq'::regclass);
ALTER TABLE ONLY bankaccount ALTER COLUMN bankaccountid SET DEFAULT nextval('bankaccount_bankaccountid_seq'::regclass);
ALTER TABLE ONLY transaction ALTER COLUMN transactionid SET DEFAULT nextval('transaction_transactionid_seq'::regclass);
ALTER TABLE ONLY transaction_type ALTER COLUMN transaction_type_id SET DEFAULT nextval('transaction_type_transaction_type_id_seq'::regclass);
ALTER TABLE ONLY users ALTER COLUMN userid SET DEFAULT nextval('users_userid_seq'::regclass);
ALTER TABLE ONLY bank
    ADD CONSTRAINT pk_bank PRIMARY KEY (bankid);
ALTER TABLE ONLY bankaccount
    ADD CONSTRAINT pk_bankaccount PRIMARY KEY (bankaccountid);

ALTER TABLE ONLY transaction_type
    ADD CONSTRAINT pk_transaction_type PRIMARY KEY (transaction_type_id);

ALTER TABLE ONLY users
    ADD CONSTRAINT pk_users PRIMARY KEY (userid);
ALTER TABLE ONLY transaction
    ADD CONSTRAINT pk_transaction PRIMARY KEY (transactionid);
ALTER TABLE ONLY bankaccount
    ADD CONSTRAINT uk_account_numbre UNIQUE (accountnumber);
ALTER TABLE ONLY transaction
    ADD CONSTRAINT fk_account FOREIGN KEY (bankaccountid) REFERENCES bankaccount(bankaccountid);
ALTER TABLE ONLY transaction
        ADD CONSTRAINT fk_type FOREIGN KEY (transaction_type_id) REFERENCES transaction_type(transaction_type_id);
ALTER TABLE ONLY bankaccount
    ADD CONSTRAINT fk_bank_account FOREIGN KEY (bankid) REFERENCES bank(bankid);
ALTER TABLE ONLY bankaccount
    ADD CONSTRAINT fk_owner_account FOREIGN KEY (userid) REFERENCES users(userid);
