//
//  CustomView1Cell.swift
//  NIOSocial_Example
//
//  Created by tianchen.xie on 2023/3/3.
//  Copyright © 2023 CocoaPods. All rights reserved.
//


import UIKit
import MPUIKit


// MARK: - reform to UITableView cell
class CustomView1Cell: UITableViewCell {
    // RENDER
    // cell.mainView.render(model)
    
    // LIFE
    override init(style: UITableViewCell.CellStyle, reuseIdentifier: String?) {
        super.init(style: style, reuseIdentifier: reuseIdentifier)
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
    
    // PROPS
    var mainView = CustomView1()
}






// MARK: - reform to UICollectionView cell
class CustomView1cCell: UICollectionViewCell {
    // RENDER
    // cell.mainView.render(model)
    
    // LIFE
    override init(frame: CGRect) {
        super.init(frame: frame)
        setup()
    }

    required init?(coder: NSCoder) {
        super.init(coder: coder)
        setup()
    }
    
    private func setup() {
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
    
    // PROPS
    var mainView = CustomView1()
}
