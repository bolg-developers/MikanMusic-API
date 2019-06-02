CREATE TABLE IF NOT EXISTS genres (
        id CHAR(36) NOT NULL,
        name CHAR(255) NOT NULL UNIQUE,
        PRIMARY KEY(id)
);

INSERT INTO genres (id,name) VALUES('a1fb538b-f081-4362-80a9-9bc8cec053db','ポップ');
INSERT INTO genres (id,name) VALUES('e88772fe-abea-4e91-894a-3b2d599ff673','チルアウト');
INSERT INTO genres (id,name) VALUES('a0a1b419-4b75-413e-b32f-742668f1e74d','エレクトリック/ダンス');
INSERT INTO genres (id,name) VALUES('0110fc9a-5c06-44cf-a0dd-97017971f145','ジャズ');
INSERT INTO genres (id,name) VALUES('34d74409-781f-4b83-8a9a-80d0d9a24f2a','ヒップホップ');
INSERT INTO genres (id,name) VALUES('468881a3-a0f5-42af-9f04-12a1c9c692d7','R&B');
INSERT INTO genres (id,name) VALUES('dfd41d95-11cf-433a-b319-bf4c74bcad5f','ロック');
INSERT INTO genres (id,name) VALUES('a1a317f8-5809-447c-b353-71c85266df0e','インディー/オルタナティブ');
INSERT INTO genres (id,name) VALUES('be3bf339-90d9-4a54-8529-97915ae91e00','クラシック');
INSERT INTO genres (id,name) VALUES('e7e49846-550b-4636-80ae-b54e87ca4403','レゲエ');
INSERT INTO genres (id,name) VALUES('a98c3ba2-7414-4ea6-b337-f8bef7ea40c5','ソウル');
INSERT INTO genres (id,name) VALUES('4a821168-e734-4a98-a7a2-5794be2428c8','メタル');
INSERT INTO genres (id,name) VALUES('3997108b-34a2-448c-842d-d3311a9713c0','ラテン');
INSERT INTO genres (id,name) VALUES('16285041-153a-49ed-98a3-3ac923215d09','ブルース');
INSERT INTO genres (id,name) VALUES('2fe67180-8596-4f2d-a07f-6e676a2649cb','ファンク');
INSERT INTO genres (id,name) VALUES('15c7d984-1065-4481-b7c0-228e833f9126','パンク');
INSERT INTO genres (id,name) VALUES('a61fb80a-55f6-489c-9e78-b76fc2531432','カントリー');
INSERT INTO genres (id,name) VALUES('ea6d9027-72fd-4286-bd4e-c024d6c6d36','フォーク/アコースティック');
INSERT INTO genres (id,name) VALUES('70dc052e-d1dd-4adb-8514-881ce301e182','インド');