use account

INSERT INTO user values(1, "testuser", "testpassword", "2019-10-13 13:45:57", "2019-10-13 13:45:57");

use resource

INSERT INTO shop values
    (1, "shop1", 0, "https://www.google.com/", 1, "2019-10-13 13:45:57", "2019-10-13 13:45:57"),
    (2, "shop2", 1, "https://www.yahoo.co.jp/", 1, "2019-10-13 13:45:57", "2019-10-13 13:45:57"),
    (3, "shop3", 2, "https://www.mercari.com/jp/", 1, "2019-10-13 13:45:57", "2019-10-13 13:45:57");

INSERT INTO tag values
    (1, "早い", 0, "2019-10-13 13:45:57", "2019-10-13 13:45:57"),
    (2, "安い", 0, "2019-10-13 13:45:57", "2019-10-13 13:45:57"),
    (3, "うまい", 0, "2019-10-13 13:45:57", "2019-10-13 13:45:57");

INSERT INTO tag_map values
    (1, 1, 1, "2019-10-13 13:45:57", "2019-10-13 13:45:57"),
    (2, 1, 2, "2019-10-13 13:45:57", "2019-10-13 13:45:57"),
    (3, 1, 3, "2019-10-13 13:45:57", "2019-10-13 13:45:57"),
    (4, 2, 2, "2019-10-13 13:45:57", "2019-10-13 13:45:57"),
    (5, 3, 3, "2019-10-13 13:45:57", "2019-10-13 13:45:57");
