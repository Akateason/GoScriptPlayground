class CustomView3: UIView {
    override init(frame: CGRect) {
        super.init(frame: frame)
        
        backgroundColor = .red
        
        let sub = UIView()
        addSubview(sub)
        sub.snp.makeConstraints { make in
            make.top.left.bottom.equalToSuperview()
            make.height.equalTo(100)
        }
        sub.backgroundColor = .green
    }
    
    required init?(coder: NSCoder) {
        fatalError("init(coder:) has not been implemented")
    }
}

