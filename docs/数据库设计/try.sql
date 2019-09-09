/*==============================================================*/
/* DBMS name:      MySQL 5.0                                    */
/* Created on:     2019/9/9 19:28:08                            */
/*==============================================================*/


drop table if exists Good;

drop table if exists Relationship_5;

drop table if exists chat_log;

drop table if exists "release";

drop table if exists tag;

drop table if exists transaction;

drop table if exists user;

/*==============================================================*/
/* Table: Good                                                  */
/*==============================================================*/
create table Good
(
   good_id              int not null,
   good_name            varchar(64) not null,
   price                decimal not null,
   description          text,
   content_id           char(24) not null,
   primary key (good_id)
);

/*==============================================================*/
/* Table: Relationship_5                                        */
/*==============================================================*/
create table Relationship_5
(
   good_id              int not null,
   tag_id               int not null,
   primary key (good_id, tag_id)
);

/*==============================================================*/
/* Table: chat_log                                              */
/*==============================================================*/
create table chat_log
(
   chat_log_id          int not null,
   user_id2             int not null,
   use_user_id2         int not null,
   chat_content_id      char(24) not null,
   primary key (chat_log_id)
);

/*==============================================================*/
/* Table: "release"                                             */
/*==============================================================*/
create table "release"
(
   release_id2          int not null,
   user_id2             int not null,
   good_id              int not null,
   transaction_id       int,
   release_time         datetime not null,
   valid_time           datetime not null,
   state                int not null,
   primary key (release_id2)
);

/*==============================================================*/
/* Table: tag                                                   */
/*==============================================================*/
create table tag
(
   tag_id               int not null,
   tag_name             varchar(16) not null,
   primary key (tag_id)
);

/*==============================================================*/
/* Table: transaction                                           */
/*==============================================================*/
create table transaction
(
   transaction_id       int not null,
   user_id2             int not null,
   create_time          datetime not null,
   primary key (transaction_id)
);

/*==============================================================*/
/* Table: user                                                  */
/*==============================================================*/
create table user
(
   user_id2             int not null,
   user_name            varchar(16) not null,
   avatar_id            char(24) not null,
   telephone            char(11),
   student_id           varchar(32) not null,
   student_name         varchar(32) not null,
   primary key (user_id2),
   key AK_Identifier_2 (student_id)
);

alter table Relationship_5 add constraint FK_Relationship_5 foreign key (good_id)
      references Good (good_id) on delete restrict on update restrict;

alter table Relationship_5 add constraint FK_Relationship_6 foreign key (tag_id)
      references tag (tag_id) on delete restrict on update restrict;

alter table chat_log add constraint FK_Relationship_8 foreign key (user_id2)
      references user (user_id2) on delete restrict on update restrict;

alter table chat_log add constraint FK_Relationship_9 foreign key (use_user_id2)
      references user (user_id2) on delete restrict on update restrict;

alter table "release" add constraint FK_Relationship_3 foreign key (transaction_id)
      references transaction (transaction_id) on delete restrict on update restrict;

alter table "release" add constraint FK_Relationship_4 foreign key (good_id)
      references Good (good_id) on delete restrict on update restrict;

alter table "release" add constraint FK_Relationship_7 foreign key (user_id2)
      references user (user_id2) on delete restrict on update restrict;

alter table transaction add constraint FK_Relationship_1 foreign key (user_id2)
      references user (user_id2) on delete restrict on update restrict;

