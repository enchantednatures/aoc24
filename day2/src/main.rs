use std::cmp::Ordering;

fn main() {
    let txt = include_str!("../example.txt");
    txt.split("\n").for_each(|line| {
        println!("{}", line);
    });

    let s = txt
        .split("\n")
        .filter(|x| !x.trim().is_empty())
        .map(proc)
        .filter(|x| *x)
        .count();

    dbg!(s);

    println!("Hello, world!");
}

fn proc(s: &str) -> bool {
    let items: Vec<_> = s
        .split(" ")
        .filter(|i| !i.trim().is_empty())
        .flat_map(|x| x.parse::<i64>())
        .collect();

    dbg!(&items);

    if (items.is_sorted_by(|x, y| x > y) || items.is_sorted_by(|x, y| x < y))
        && items.windows(2).all(|x| x[0].abs_diff(x[1]) <= 3)
    {
        return true;
    }

    items.iter().enumerate().any(|(idx, item)| {
        let mut items = items.clone();
        items.remove(idx);
        if (items.is_sorted_by(|x, y| x > y) || items.is_sorted_by(|x, y| x < y))
            && items.windows(2).all(|x| x[0].abs_diff(x[1]) <= 3)
        {
            return true;
        }

        return false;
    })

    // for (item, idx) in items{
    //     // let items = items[]

    // }
}
