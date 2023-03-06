import UIKit


// table cell
class CustomViewCell: UITableViewCell {
// RENDER .
// cell.mainView.render(model)

  override init (style: UITableViewCell.CellStyle, reuseIdentifier: String?) {
      super.init (style: style, reuseIdentifier: reuseIdentifier) 
      setup()    
  }

  required init?(coder: NSCoder) {
      super.init(coder: coder)
      setup()
  }

  func setup() {
      contentView.addSubview(mainView)
      mainView.snp.makeConstraints { make in
        make.edges.equalToSuperview()
      }

      // setup mainView ...
  }

  var mainView = CustomView()
}



// collection cell
class CustomViewcCell: UICollectionViewCell {
// RENDER .
// cell.mainView.render(model)

  override init (frame: CGRect) {
      super.init(frame: frame)
      setup()
  }

  required init?(coder: NSCoder) {
      super.init(coder: coder)
      setup()
  }

  func setup() {
      contentView.translatesAutoresizingMaskIntoConstraints = false
      contentView.widthAnchor.constraint(equalToConstant: UIScreen.main.bounds.width).isActive = true

      contentView.addSubview(mainView)
      mainView.snp.makeConstraints { make in
          make.edges.equalToSuperview()
      }

      // setup mainView ...
  }

  override func preferredLayoutAttributesFitting(_ layoutAttributes: UICollectionViewLayoutAttributes) -> UICollectionViewLayoutAttributes {
      let size = self.contentView.systemLayoutSizeFitting(layoutAttributes.size)
      var cellFrame = layoutAttributes.frame
      cellFrame.size.height = size.height
      cellFrame.size.width = UIScreen.main.bounds.width
      layoutAttributes.frame = cellFrame
      return layoutAttributes
  }

  var mainView = CustomView()
}