// クラス

class Point {
  x: number;
  y: number;
  constructor(x: number, y: number) {
      this.x = x;
      this.y = y;
  }
  add(point: Point) {
      return new Point(this.x + point.x, this.y + point.y);
  }
}

var p1 = new Point(0, 10);
var p2 = new Point(10, 20);
var p3 = p1.add(p2);

console.log(p1);
console.log(p2);
console.log(p3);

// 出力
// Point { x: 0, y: 10 }
// Point { x: 10, y: 20 }
// Point { x: 10, y: 30 }
