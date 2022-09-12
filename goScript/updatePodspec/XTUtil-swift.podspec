Pod::Spec.new do |s|
  s.name             = 'XTUtil-swift'
	s.version = '4.0.4'
  s.summary          = 'A short description of XTUtil-swift.'
  s.description      = <<-DESC
TODO: Add long description of the pod here.
                       DESC
  s.homepage         = 'https://github.com/akateason/XTUtil-swift'
  s.license          = { :type => 'MIT', :file => 'LICENSE' }
  s.author           = { 'teason' => 'akateason@qq.com' }
  s.source           = { :git => 'https://gitee.com/mamba24xtc/XTUtil-swift', :tag => s.version.to_s }

  s.ios.deployment_target = '9.0'

  s.source_files = 'XTUtil-swift/Source/**/*'

  # s.resource_bundles = {
  #   'XTUtil-swift' => ['XTUtil-swift/Assets/*.png']
  # }

  #s.subspec 'XTUtil-swift' do | sm |
      #sm.source_files = 'XTUtil-swift/ZYSubModule/**/*'
      #sm.dependency 'AFNetworking', '~> 2.3'
  #end

  # s.public_header_files = 'Pod/Classes/**/*.h'
  # s.frameworks = 'UIKit', 'MapKit'
  # s.dependency 'AFNetworking', '~> 2.3'

  s.dependency 'SnapKit'
  s.dependency 'HandyJSON'
  s.dependency 'RxSwift'
  s.dependency 'RxGesture'
  s.dependency 'SmoothFrame','1.0.0.1'

end
